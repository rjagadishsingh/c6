// generated by stringer -type=TokenType token.go; DO NOT EDIT

package ast

import "fmt"

const _TokenType_name = "T_SPACET_COMMENT_LINET_COMMENT_BLOCKT_SEMICOLONT_COMMAT_IDENTT_URLT_MEDIAT_TRUET_FALSET_NULLT_MS_PARAM_NAMET_FUNCTION_NAMET_ID_SELECTORT_CLASS_SELECTORT_TYPE_SELECTORT_UNIVERSAL_SELECTORT_PARENT_SELECTORT_PSEUDO_SELECTORT_INTERPOLATION_SELECTORT_LITERAL_CONCATT_MS_PROGIDT_AND_SELECTORT_DESCENDANT_COMBINATORT_CHILD_COMBINATORT_ADJACENT_SIBLING_COMBINATORT_UNICODE_RANGET_IFT_ELSET_ELSE_IFT_INCLUDET_MIXINT_FUNCTIONT_GLOBALT_DEFAULTT_IMPORTANTT_OPTIONALT_FONT_FACET_ORT_ANDT_XORT_PLUST_BRACE_STARTT_BRACE_ENDT_LANG_CODET_BRACKET_LEFTT_ATTRIBUTE_NAMET_BRACKET_RIGHTT_EQUALT_GTT_GET_LTT_LET_ATTR_EQUALT_ATTR_TILDE_EQUALT_ATTR_HYPHEN_EQUALT_VARIABLET_IMPORTT_AT_RULET_CHARSETT_QQ_STRINGT_Q_STRINGT_UNQUOTE_STRINGT_PAREN_STARTT_PAREN_ENDT_CONSTANTT_INTEGERT_FLOATT_UNIT_PERCENTT_UNIT_SECONDT_UNIT_MILLISECONDT_UNIT_CHT_UNIT_CMT_UNIT_EMT_UNIT_EXT_UNIT_INT_UNIT_MMT_UNIT_PCT_UNIT_PTT_UNIT_PXT_UNIT_REMT_UNIT_HZT_UNIT_KHZT_UNIT_DPIT_UNIT_DPCMT_UNIT_DPPXT_UNIT_VHT_UNIT_VWT_UNIT_VMINT_UNIT_VMAXT_UNIT_DEGT_UNIT_GRADT_UNIT_RADT_UNIT_TURNT_PROPERTY_NAME_TOKENT_PROPERTY_VALUET_HEX_COLORT_COLONT_INTERPOLATION_STARTT_INTERPOLATION_INNERT_INTERPOLATION_ENDT_DIVT_MULT_MINUS"

var _TokenType_index = [...]uint16{0, 7, 21, 36, 47, 54, 61, 66, 73, 79, 86, 92, 107, 122, 135, 151, 166, 186, 203, 220, 244, 260, 271, 285, 306, 322, 341, 356, 360, 366, 375, 384, 391, 401, 409, 418, 429, 439, 450, 454, 459, 464, 470, 483, 494, 505, 519, 535, 550, 557, 561, 565, 569, 573, 585, 603, 622, 632, 640, 649, 658, 669, 679, 695, 708, 719, 729, 738, 745, 759, 772, 790, 799, 808, 817, 826, 835, 844, 853, 862, 871, 881, 890, 900, 910, 921, 932, 941, 950, 961, 972, 982, 993, 1003, 1014, 1035, 1051, 1062, 1069, 1090, 1111, 1130, 1135, 1140, 1147}

func (i TokenType) String() string {
	if i < 0 || i+1 >= TokenType(len(_TokenType_index)) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
