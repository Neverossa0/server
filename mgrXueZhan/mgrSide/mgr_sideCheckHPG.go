package mgrSide

import (
	"server/mgrXueZhan/mgrCard"
	"server/pb"

	"github.com/name5566/leaf/log"
)

func (sider *SideInfo) CheckHPGByOwn() []pb.ProcType {

	procList := make([]pb.ProcType, 0)
	handList := mgrCard.TransToIntList(sider.GethandCards())
	gangList := mgrCard.TransToIntList(sider.GetgangCards())
	pengList := mgrCard.TransToIntList(sider.GetpengCards())
	log.Debug("sidercheckhpg==>玩家[%v]准备摸牌后进行胡碰检测，手牌列表=[%v]", sider.GetPlayer().GetPlayerId(), handList)

	if !sider.checkhandCardsLackType() {
		if mgrCard.CheckHuOwn(handList, gangList, pengList) {
			procList = append(procList, pb.ProcType_Proc_Hu)
			sider.SetProcess(ProcessStatus_WAITING_HU)
		}
	}
	gangCard, ok := mgrCard.CheckGangOwn(handList)
	if ok {
		if !sider.checkCardLackType(int32(gangCard)) {
			procList = append(procList, pb.ProcType_Proc_Gang)
		}
	}

	if len(procList) == 0 {
		procList = append(procList, pb.ProcType_Proc_Discard)
	} else {
		procList = append(procList, pb.ProcType_Proc_Pass)
	}

	log.Debug("sidercheckhpg==>玩家[%v]对自己摸牌检测胡碰杠完毕,可进行操作=[%v]", sider.GetPlayer().GetPlayerId(), procList)
	return procList
}

func (sider *SideInfo) CheckHPGByOther(drawcard *mgrCard.CardInfo) []pb.ProcType {

	procList := make([]pb.ProcType, 0)

	if !sider.checkCardLackType(drawcard.Id) {
		handList := mgrCard.TransToIntList(sider.GethandCards())
		gangList := mgrCard.TransToIntList(sider.GetgangCards())
		pengList := mgrCard.TransToIntList(sider.GetpengCards())
		log.Debug("sidercheckhpg==>玩家[%v]准备对他人出牌进行胡碰检测，手牌列表=[%v]", sider.GetPlayer().GetPlayerId(), handList)

		if !sider.checkhandCardsLackType() {
			if mgrCard.CheckHuOther(handList, gangList, pengList, int(drawcard.Id)) {
				procList = append(procList, pb.ProcType_Proc_Hu)
				sider.SetProcess(ProcessStatus_WAITING_HU)
			}
		}

		if mgrCard.CheckGangOther(handList, drawcard) {
			procList = append(procList, pb.ProcType_Proc_Gang)
		}
		if mgrCard.CheckPengOther(handList, drawcard) {
			procList = append(procList, pb.ProcType_Proc_Peng)
		}

	}

	procList = append(procList, pb.ProcType_Proc_Pass)
	log.Debug("sidercheckhpg==>玩家[%v]对其他玩家出牌[%v]胡碰杠完毕,可进行操作=[%v]", sider.GetPlayer().GetPlayerId(), drawcard.Id, procList)

	if len(procList) == 1 {
		sider.SetProcess(ProcessStatus_TURN_OVER)
	}
	return procList
}

//============================****发送玩家可操作行为****============================
func (sider *SideInfo) SendActionListByDeal(procList []pb.ProcType) {
	data := &pb.GS2CInterruptAction{
		ProcList: procList,
	}
	sider.agent.WriteMsg(data)
}

func (sider *SideInfo) SendActionListByDiscard(procList []pb.ProcType, drawcard *mgrCard.CardInfo) {
	data := &pb.GS2CInterruptAction{
		ProcList: procList,
		DrawCard: drawcard.ToPBCard(),
	}
	sider.agent.WriteMsg(data)
}
