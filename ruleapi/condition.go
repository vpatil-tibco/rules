package ruleapi

import "github.com/TIBCOSoftware/bego/common/model"

//condition ... is a rete condtion
type condition interface {
	GetName() string
	GetEvaluator() model.ConditionEvaluator
	GetRule() Rule
	GetStreamSource() []model.StreamSource
	//Stringer.String interface
	String() string
}

type conditionImpl struct {
	name        string
	rule        Rule
	identifiers []model.StreamSource
	cfn         model.ConditionEvaluator
}

//NewCondition ... a new Condition
func NewCondition(name string, rule Rule, identifiers []model.StreamSource, cfn model.ConditionEvaluator) condition {
	c := conditionImpl{}
	c.initConditionImpl(name, rule, identifiers, cfn)
	return &c
}

func (cnd *conditionImpl) initConditionImpl(name string, rule Rule, identifiers []model.StreamSource, cfn model.ConditionEvaluator) {
	cnd.name = name
	cnd.rule = rule
	cnd.identifiers = append(cnd.identifiers, identifiers...)
	cnd.cfn = cfn
}

func (cnd *conditionImpl) GetIdentifiers() []model.StreamSource {
	return cnd.identifiers
}

func (cnd *conditionImpl) GetEvaluator() model.ConditionEvaluator {
	return cnd.cfn
}

func (cnd *conditionImpl) String() string {
	return "[Condition: name:" + cnd.name + ", idrs: TODO]"
}

func (cnd *conditionImpl) GetName() string {
	return cnd.name
}

func (cnd *conditionImpl) GetRule() Rule {
	return cnd.rule
}
func (cnd *conditionImpl) GetStreamSource() []model.StreamSource {
	return cnd.identifiers
}