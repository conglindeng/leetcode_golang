package easy

func MinNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	experienceNeed:=0
	energyNeed:=0
	n:=len(experience)
	curEnergy:=initialEnergy
	curExperience:=initialExperience
	for i:=0;i<n;i++{
		ener:=energy[i]
		exper:=experience[i]
		if curEnergy<=ener{
			energyNeed+=ener-curEnergy+1
			curEnergy=1
		}else{
			curEnergy-=ener
		}
		if curExperience<=exper{
			experienceNeed+=exper-curExperience+1
			curExperience=exper*2+1
		}else{
			curExperience+=exper
		}
	}
	return energyNeed+experienceNeed
}