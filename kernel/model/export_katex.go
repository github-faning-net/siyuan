// SiYuan - Refactor your thinking
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package model

import (
	"sort"
	"strconv"
	"strings"

	"github.com/88250/gulu"
)

func extractUsedMacros(mathContent string, macrosKeys *[]string) (ret []string) {
Next:
	for {
		for _, k := range *macrosKeys {
			if idx := strings.Index(mathContent, k); -1 < idx {
				mathContent = strings.Replace(mathContent, k, "__@"+k[1:]+"@__", 1)
				ret = append(ret, k)
				goto Next
			}
		}
		break
	}
	return
}

var katexSupportedFunctions = []string{ // https://katex.org/docs/supported.html

	"\\!",
	"\\\"",
	"\\#",
	"\\$",
	"\\%",
	"\\&",
	"\\'",
	"\\,",
	"\\.",
	"\\:",
	"\\;",
	"\\=",
	"\\>",
	"\\Alpha",
	"\\And",
	"\\Bbb",
	"\\Bbbk",
	"\\Beta",
	"\\Big",
	"\\Bigg",
	"\\Biggl",
	"\\Biggm",
	"\\Biggr",
	"\\Bigl",
	"\\Bigm",
	"\\Bigr",
	"\\Box",
	"\\Bra",
	"\\Braket",
	"\\Bumpeq",
	"\\Cap",
	"\\Chi",
	"\\Colonapprox",
	"\\Coloneq",
	"\\Coloneqq",
	"\\Colonsim",
	"\\Complex",
	"\\Cup",
	"\\Dagger",
	"\\Darr",
	"\\Delta",
	"\\Diamond",
	"\\Doteq",
	"\\Downarrow",
	"\\Epsilon",
	"\\Eqcolon",
	"\\Eqqcolon",
	"\\Eta",
	"\\Finv",
	"\\Game",
	"\\Gamma",
	"\\H",
	"\\Harr",
	"\\Huge AB",
	"\\Im",
	"\\Iota",
	"\\Join",
	"\\KaTeX",
	"\\Kappa",
	"\\Ket",
	"\\LARGE AB",
	"\\LaTeX",
	"\\Lambda",
	"\\Large AB",
	"\\Larr",
	"\\Leftarrow",
	"\\Leftrightarrow",
	"\\Lleftarrow",
	"\\Longleftarrow",
	"\\Longleftrightarrow",
	"\\Longrightarrow",
	"\\Lrarr",
	"\\Lsh",
	"\\Mu",
	"\\N",
	"\\Nu",
	"\\Omega",
	"\\Omicron",
	"\\Overrightarrow",
	"\\P",
	"\\Phi",
	"\\Pi",
	"\\Pr",
	"\\Psi",
	"\\R",
	"\\Rarr",
	"\\Re",
	"\\Reals",
	"\\Rho",
	"\\Rightarrow",
	"\\Rrightarrow",
	"\\Rsh",
	"\\S",
	"\\Set",
	"\\Sigma",
	"\\Subset",
	"\\Supset",
	"\\Tau",
	"\\TeX",
	"\\Theta",
	"\\Uarr",
	"\\Uparrow",
	"\\Updownarrow",
	"\\Upsilon",
	"\\Vdash",
	"\\Vert",
	"\\Vvdash",
	"\\Xi",
	"\\Z",
	"\\Zeta",
	"\\^",
	"\\_",
	"\\`",
	"\\acute",
	"\\alef",
	"\\alefsym",
	"\\aleph",
	"\\alpha",
	"\\amalg",
	"\\angle",
	"\\approx",
	"\\approxcolon",
	"\\approxcoloncolon",
	"\\approxeq",
	"\\arccos",
	"\\arcctg",
	"\\arcsin",
	"\\arctan",
	"\\arctg",
	"\\arg",
	"\\argmax",
	"\\argmin",
	"\\ast",
	"\\asymp",
	"\\backepsilon",
	"\\backprime",
	"\\backsim",
	"\\backsimeq",
	"\\backslash",
	"\\bar",
	"\\barwedge",
	"\\bcancel",
	"\\because",
	"\\begin",
	"\\beta",
	"\\beth",
	"\\between",
	"\\bf Ab0",
	"\\big",
	"\\bigcap",
	"\\bigcirc",
	"\\bigcup",
	"\\bigg",
	"\\biggl",
	"\\biggm",
	"\\biggr",
	"\\bigl",
	"\\bigm",
	"\\bigodot",
	"\\bigoplus",
	"\\bigotimes",
	"\\bigr",
	"\\bigsqcup",
	"\\bigstar",
	"\\bigtriangledown",
	"\\bigtriangleup",
	"\\biguplus",
	"\\bigvee",
	"\\bigwedge",
	"\\binom",
	"\\blacklozenge",
	"\\blacksquare",
	"\\blacktriangle",
	"\\blacktriangledown",
	"\\blacktriangleleft",
	"\\blacktriangleright",
	"\\bm",
	"\\bmod",
	"\\bold",
	"\\boldsymbol",
	"\\bot",
	"\\bowtie",
	"\\boxdot",
	"\\boxed",
	"\\boxminus",
	"\\boxplus",
	"\\boxtimes",
	"\\bra",
	"\\braket",
	"\\breve",
	"\\bullet",
	"\\bumpeq",
	"\\cal AB0",
	"\\cancel",
	"\\cap",
	"\\cdot",
	"\\cdotp",
	"\\cdots",
	"\\centerdot",
	"\\cfrac",
	"\\ch",
	"\\check",
	"\\checkmark",
	"\\chi",
	"\\circ",
	"\\circeq",
	"\\circlearrowleft",
	"\\circlearrowright",
	"\\circledR",
	"\\circledS",
	"\\circledast",
	"\\circledcirc",
	"\\circleddash",
	"\\clubs",
	"\\clubsuit",
	"\\cnums",
	"\\colon",
	"\\colonapprox",
	"\\coloncolon",
	"\\coloncolonapprox",
	"\\coloncolonequals",
	"\\coloncolonminus",
	"\\coloncolonsim",
	"\\coloneq",
	"\\coloneqq",
	"\\colonequals",
	"\\colonminus",
	"\\colonsim",
	"\\complement",
	"\\cong",
	"\\coprod",
	"\\copyright",
	"\\cos",
	"\\cosec",
	"\\cosh",
	"\\cot",
	"\\cotg",
	"\\coth",
	"\\csc",
	"\\ctg",
	"\\cth",
	"\\cup",
	"\\curlyeqprec",
	"\\curlyeqsucc",
	"\\curlyvee",
	"\\curlywedge",
	"\\curvearrowleft",
	"\\curvearrowright",
	"\\dArr",
	"\\dag",
	"\\dagger",
	"\\daleth",
	"\\darr",
	"\\dashleftarrow",
	"\\dashrightarrow",
	"\\dashv",
	"\\dbinom",
	"\\dblcolon",
	"\\ddag",
	"\\ddagger",
	"\\ddot",
	"\\ddots",
	"\\def\\arraystretch",
	"\\def\\foo",
	"\\deg",
	"\\degree",
	"\\delta",
	"\\det",
	"\\dfrac",
	"\\diagdown",
	"\\diagup",
	"\\diamond",
	"\\diamonds",
	"\\diamondsuit",
	"\\digamma",
	"\\dim",
	"\\displaystyle\\sum_",
	"\\div",
	"\\divideontimes",
	"\\dot",
	"\\doteq",
	"\\doteqdot",
	"\\dotplus",
	"\\dots",
	"\\dotsb",
	"\\dotsc",
	"\\dotsi",
	"\\dotsm",
	"\\dotso",
	"\\doublebarwedge",
	"\\doublecap",
	"\\doublecup",
	"\\downarrow",
	"\\downdownarrows",
	"\\downharpoonleft",
	"\\downharpoonright",
	"\\edef\\macroname#1#2…",
	"\\ell",
	"\\empty",
	"\\emptyset",
	"\\end",
	"\\enspace",
	"\\epsilon",
	"\\eqcirc",
	"\\eqcolon",
	"\\eqqcolon",
	"\\eqsim",
	"\\eqslantgtr",
	"\\eqslantless",
	"\\equalscolon",
	"\\equalscoloncolon",
	"\\equiv",
	"\\eta",
	"\\eth",
	"\\exist",
	"\\exists",
	"\\exp",
	"\\fallingdotseq",
	"\\flat",
	"\\footnotesize AB",
	"\\forall",
	"\\frac",
	"\\frak",
	"\\frown",
	"\\futurelet\\foo\\bar x",
	"\\gamma",
	"\\gcd",
	"\\gdef\\bar#1",
	"\\ge",
	"\\genfrac ( ] ",
	"\\geq",
	"\\geqq",
	"\\geqslant",
	"\\gets",
	"\\gg",
	"\\ggg",
	"\\gggtr",
	"\\gimel",
	"\\global\\def\\macroname#1#2…",
	"\\gnapprox",
	"\\gneq",
	"\\gneqq",
	"\\gnsim",
	"\\grave",
	"\\gt",
	"\\gtrapprox",
	"\\gtrdot",
	"\\gtreqless",
	"\\gtreqqless",
	"\\gtrless",
	"\\gtrsim",
	"\\gvertneqq",
	"\\hArr",
	"\\harr",
	"\\hat",
	"\\hbar",
	"\\hdashline",
	"\\hearts",
	"\\heartsuit",
	"\\hom",
	"\\hookleftarrow",
	"\\hookrightarrow",
	"\\hphantom",
	"\\href",
	"\\hskip",
	"\\hslash",
	"\\hspace",
	"\\hspace*",
	"\\htmlClass",
	"\\htmlData",
	"\\htmlId",
	"\\htmlStyle",
	"\\huge AB",
	"\\iff",
	"\\iiint",
	"\\iint",
	"\\image",
	"\\imageof",
	"\\imath",
	"\\impliedby",
	"\\implies",
	"\\in",
	"\\includegraphics[height=0.8em, totalheight=0.9em, width=0.9em, alt=KA logo]",
	"\\inf",
	"\\infin",
	"\\infty",
	"\\injlim",
	"\\int",
	"\\intercal",
	"\\intop",
	"\\iota",
	"\\isin",
	"\\it Ab0",
	"\\jmath",
	"\\kappa",
	"\\ker",
	"\\kern",
	"\\ket",
	"\\lArr",
	"\\lBrace \\rBrace",
	"\\lVert",
	"\\lambda",
	"\\land",
	"\\lang",
	"\\langle",
	"\\large AB",
	"\\larr",
	"\\lbrace",
	"\\lbrack",
	"\\lceil",
	"\\ldotp",
	"\\ldots",
	"\\le",
	"\\leadsto",
	"\\left",
	"\\left(x^",
	"\\left.",
	"\\leftarrow",
	"\\leftarrowtail",
	"\\leftharpoondown",
	"\\leftharpoonup",
	"\\leftleftarrows",
	"\\leftrightarrow",
	"\\leftrightarrows",
	"\\leftrightharpoons",
	"\\leftrightsquigarrow",
	"\\leftthreetimes",
	"\\leq",
	"\\leqq",
	"\\leqslant",
	"\\lessapprox",
	"\\lessdot",
	"\\lesseqgtr",
	"\\lesseqqgtr",
	"\\lessgtr",
	"\\lesssim",
	"\\let\\foo=\\bar",
	"\\lfloor",
	"\\lg",
	"\\lgroup",
	"\\lhd",
	"\\lim",
	"\\lim\\limits_x",
	"\\lim\\nolimits_x",
	"\\liminf",
	"\\limsup",
	"\\ll",
	"\\llbracket",
	"\\llcorner",
	"\\lll",
	"\\llless",
	"\\lmoustache",
	"\\ln",
	"\\lnapprox",
	"\\lneq",
	"\\lneqq",
	"\\lnot",
	"\\lnsim",
	"\\log",
	"\\longleftarrow",
	"\\longleftrightarrow",
	"\\longmapsto",
	"\\longrightarrow",
	"\\looparrowleft",
	"\\looparrowright",
	"\\lor",
	"\\lozenge",
	"\\lparen",
	"\\lq",
	"\\lrArr",
	"\\lrarr",
	"\\lrcorner",
	"\\lt",
	"\\lt \\gt",
	"\\ltimes",
	"\\lvert",
	"\\lvertneqq",
	"\\maltese",
	"\\mapsto",
	"\\mathbb",
	"\\mathbf",
	"\\mathcal",
	"\\mathellipsis",
	"\\mathfrak",
	"\\mathit",
	"\\mathnormal",
	"\\mathring",
	"\\mathrlap",
	"\\mathrm",
	"\\mathscr",
	"\\mathsf",
	"\\mathsterling",
	"\\mathstrut",
	"\\mathtt",
	"\\max",
	"\\measuredangle",
	"\\medspace",
	"\\mho",
	"\\mid",
	"\\middle",
	"\\min",
	"\\minuscolon",
	"\\minuscoloncolon",
	"\\minuso",
	"\\mkern",
	"\\models",
	"\\mp",
	"\\mskip",
	"\\mu",
	"\\multimap",
	"\\nLeftarrow",
	"\\nLeftrightarrow",
	"\\nRightarrow",
	"\\nVDash",
	"\\nVdash",
	"\\nabla",
	"\\natnums",
	"\\natural",
	"\\ncong",
	"\\ne",
	"\\nearrow",
	"\\neg",
	"\\negmedspace",
	"\\negthickspace",
	"\\negthinspace",
	"\\neq",
	"\\newcommand\\macroname[numargs]",
	"\\nexists",
	"\\ngeq",
	"\\ngeqq",
	"\\ngeqslant",
	"\\ngtr",
	"\\ni",
	"\\nleftarrow",
	"\\nleftrightarrow",
	"\\nleq",
	"\\nleqq",
	"\\nleqslant",
	"\\nless",
	"\\nmid",
	"\\nobreakspace",
	"\\nonumber",
	"\\normalsize AB",
	"\\not =",
	"\\notag",
	"\\notin",
	"\\notni",
	"\\nparallel",
	"\\nprec",
	"\\npreceq",
	"\\nrightarrow",
	"\\nshortmid",
	"\\nshortparallel",
	"\\nsim",
	"\\nsubseteq",
	"\\nsubseteqq",
	"\\nsucc",
	"\\nsucceq",
	"\\nsupseteq",
	"\\nsupseteqq",
	"\\ntriangleleft",
	"\\ntrianglelefteq",
	"\\ntriangleright",
	"\\ntrianglerighteq",
	"\\nu",
	"\\nvDash",
	"\\nvdash",
	"\\nwarrow",
	"\\odot",
	"\\oiiint",
	"\\oiint",
	"\\oint",
	"\\omega",
	"\\omicron",
	"\\ominus",
	"\\operatorname",
	"\\operatorname*",
	"\\operatornamewithlimits",
	"\\oplus",
	"\\origof",
	"\\oslash",
	"\\otimes",
	"\\overbrace",
	"\\overgroup",
	"\\overleftarrow",
	"\\overleftharpoon",
	"\\overleftrightarrow",
	"\\overline",
	"\\overlinesegment",
	"\\overrightarrow",
	"\\overrightharpoon",
	"\\overset",
	"\\owns",
	"\\parallel",
	"\\partial",
	"\\perp",
	"\\phantom",
	"\\phase",
	"\\phi",
	"\\pi",
	"\\pitchfork",
	"\\plim",
	"\\plusmn",
	"\\pm",
	"\\pounds",
	"\\prec",
	"\\precapprox",
	"\\preccurlyeq",
	"\\preceq",
	"\\precnapprox",
	"\\precneqq",
	"\\precnsim",
	"\\precsim",
	"\\prime",
	"\\prod",
	"\\projlim",
	"\\propto",
	"\\providecommand\\macroname[numargs]",
	"\\psi",
	"\\qquad",
	"\\quad",
	"\\r",
	"\\rArr",
	"\\rVert",
	"\\rang",
	"\\rangle",
	"\\rarr",
	"\\ratio",
	"\\rbrace",
	"\\rbrack",
	"\\rceil",
	"\\real",
	"\\reals",
	"\\renewcommand\\macroname[numargs]",
	"\\restriction",
	"\\rfloor",
	"\\rgroup",
	"\\rhd",
	"\\rho",
	"\\right",
	"\\right.",
	"\\rightarrow",
	"\\rightarrowtail",
	"\\rightharpoondown",
	"\\rightharpoonup",
	"\\rightleftarrows",
	"\\rightleftharpoons",
	"\\rightrightarrows",
	"\\rightsquigarrow",
	"\\rightthreetimes",
	"\\risingdotseq",
	"\\rm Ab0",
	"\\rmoustache",
	"\\rparen",
	"\\rq",
	"\\rrbracket",
	"\\rtimes",
	"\\rvert",
	"\\scriptscriptstyle x",
	"\\scriptsize AB",
	"\\scriptstyle x",
	"\\sdot",
	"\\searrow",
	"\\sec",
	"\\set",
	"\\setminus",
	"\\sf Ab0",
	"\\sh",
	"\\sharp",
	"\\shortmid",
	"\\shortparallel",
	"\\sigma",
	"\\sim",
	"\\simcolon",
	"\\simcoloncolon",
	"\\simeq",
	"\\sin",
	"\\sinh",
	"\\small AB",
	"\\smallfrown",
	"\\smallint",
	"\\smallsetminus",
	"\\smallsmile",
	"\\smile",
	"\\sout",
	"\\space",
	"\\spades",
	"\\spadesuit",
	"\\sphericalangle",
	"\\sqcap",
	"\\sqcup",
	"\\sqrt",
	"\\sqsubset",
	"\\sqsubseteq",
	"\\sqsupset",
	"\\sqsupseteq",
	"\\square",
	"\\stackrel",
	"\\star",
	"\\sub",
	"\\sube",
	"\\subset",
	"\\subseteq",
	"\\subseteqq",
	"\\subsetneq",
	"\\subsetneqq",
	"\\succ",
	"\\succapprox",
	"\\succcurlyeq",
	"\\succeq",
	"\\succnapprox",
	"\\succneqq",
	"\\succnsim",
	"\\succsim",
	"\\sum",
	"\\sum_",
	"\\sup",
	"\\supe",
	"\\supset",
	"\\supseteq",
	"\\supseteqq",
	"\\supsetneq",
	"\\supsetneqq",
	"\\surd",
	"\\swarrow",
	"\\tan",
	"\\tanh",
	"\\tau",
	"\\tbinom",
	"\\text",
	"\\textbf",
	"\\textit",
	"\\textmd",
	"\\textnormal",
	"\\textrm",
	"\\textsf",
	"\\textstyle\\sum_",
	"\\texttt",
	"\\textup",
	"\\tfrac",
	"\\tg",
	"\\th",
	"\\therefore",
	"\\theta",
	"\\thetasym",
	"\\thickapprox",
	"\\thicksim",
	"\\thickspace",
	"\\thinspace",
	"\\tilde",
	"\\times",
	"\\tiny AB",
	"\\to",
	"\\top",
	"\\triangle",
	"\\triangledown",
	"\\triangleleft",
	"\\trianglelefteq",
	"\\triangleq",
	"\\triangleright",
	"\\trianglerighteq",
	"\\tt Ab0",
	"\\twoheadleftarrow",
	"\\twoheadrightarrow",
	"\\u",
	"\\uArr",
	"\\uarr",
	"\\ulcorner",
	"\\underbar",
	"\\underbrace",
	"\\undergroup",
	"\\underleftarrow",
	"\\underleftrightarrow",
	"\\underline",
	"\\underlinesegment",
	"\\underrightarrow",
	"\\underset",
	"\\unlhd",
	"\\unrhd",
	"\\uparrow",
	"\\updownarrow",
	"\\upharpoonleft",
	"\\upharpoonright",
	"\\uplus",
	"\\upsilon",
	"\\upuparrows",
	"\\urcorner",
	"\\url",
	"\\utilde",
	"\\v",
	"\\vDash",
	"\\varDelta",
	"\\varGamma",
	"\\varLambda",
	"\\varOmega",
	"\\varPhi",
	"\\varPi",
	"\\varPsi",
	"\\varSigma",
	"\\varTheta",
	"\\varUpsilon",
	"\\varXi",
	"\\varepsilon",
	"\\varinjlim",
	"\\varkappa",
	"\\varliminf",
	"\\varlimsup",
	"\\varnothing",
	"\\varphi",
	"\\varpi",
	"\\varprojlim",
	"\\varpropto",
	"\\varrho",
	"\\varsigma",
	"\\varsubsetneq",
	"\\varsubsetneqq",
	"\\varsupsetneq",
	"\\varsupsetneqq",
	"\\vartheta",
	"\\vartriangle",
	"\\vartriangleleft",
	"\\vartriangleright",
	"\\vcentcolon",
	"\\vdash",
	"\\vdots",
	"\\vec",
	"\\vee",
	"\\veebar",
	"\\verb!x^2!",
	"\\vert",
	"\\vphantom",
	"\\wedge",
	"\\weierp",
	"\\widecheck",
	"\\widehat",
	"\\widetilde",
	"\\wp",
	"\\wr",
	"\\xLeftarrow",
	"\\xLeftrightarrow",
	"\\xRightarrow",
	"\\xcancel",
	"\\xdef\\macroname#1#2…",
	"\\xhookleftarrow",
	"\\xhookrightarrow",
	"\\xi",
	"\\xleftarrow",
	"\\xleftharpoondown",
	"\\xleftharpoonup",
	"\\xleftrightarrow",
	"\\xleftrightharpoons",
	"\\xlongequal",
	"\\xmapsto",
	"\\xrightarrow[under]",
	"\\xrightharpoondown",
	"\\xrightharpoonup",
	"\\xrightleftharpoons",
	"\\xtofrom",
	"\\xtwoheadleftarrow",
	"\\xtwoheadrightarrow",
	"\\yen",
	"\\zeta",
	"\\|",
	"\\~",
}

func init() {
	sort.Slice(katexSupportedFunctions, func(i, j int) bool { return len(katexSupportedFunctions[i]) > len(katexSupportedFunctions[j]) })
}

func resolveKaTexMacro(macroName string, macros *map[string]string, keys *[]string, depth *int) string {
	v := (*macros)[macroName]
	if *depth > 16 {
		// Limit KaTex macro maximum recursive parsing depth is 16 https://github.com/siyuan-note/siyuan/issues/10484
		return v
	}

	sort.Slice(*keys, func(i, j int) bool { return len((*keys)[i]) > len((*keys)[j]) })
	*depth++
	for _, k := range *keys {
		escaped := escapeKaTexSupportedFunctions(v)
		if strings.Contains(escaped, k) {
			escaped = strings.ReplaceAll(escaped, k, resolveKaTexMacro(k, macros, keys, depth))
			v = unescapeKaTexSupportedFunctions(escaped)
			(*macros)[macroName] = v
		}
	}
	return v
}

func escapeKaTexSupportedFunctions(macroVal string) string {
Next:
	for {
		for _, f := range katexSupportedFunctions {
			if idx := strings.Index(macroVal, f); -1 < idx {
				macroVal = strings.Replace(macroVal, f, "__@"+f[1:]+"@__", 1)
				goto Next
			}
		}
		break
	}
	return macroVal
}

func unescapeKaTexSupportedFunctions(macroVal string) string {
	if !strings.Contains(macroVal, "__@") {
		return macroVal
	}

	for _, f := range katexSupportedFunctions {
		macroVal = strings.ReplaceAll(macroVal, "__@"+f[1:]+"@__", f)
	}
	return macroVal
}

func fillKaTexMacrosParams(macro string, mathContent, expanded *string) {
	// Support KaTex macro parameters https://github.com/siyuan-note/siyuan/issues/11448

	idx := strings.Index(*mathContent, macro)
	if idx < 0 {
		return
	}

	// 提取 mathContent 中 {} 包裹的实参，这里可能会提取到其他多余的参数，但是后面仅替换宏中的 #1, #2, ...，所以不影响
	tmp := (*mathContent)[idx:]
	args := map[int]string{}
	for i, arg := range gulu.Str.SubstringsBetween(tmp, "{", "}") {
		args[i+1] = arg
	}

	// 将宏展开中的 #1, #2, ... 替换为实参
	// Macros accept up to nine arguments: #1, #2, etc. https://katex.org/docs/supported#macros
	paramsCount := 0
	for i := 0; i < len(args); i++ {
		if strings.Contains(*expanded, "#"+strconv.Itoa(i+1)) {
			paramsCount++
			*expanded = strings.ReplaceAll(*expanded, "#"+strconv.Itoa(i+1), args[i+1])
		}
	}

	// 根据参数个数将 mathContent 中的实参替换为空
	for i := 1; i <= paramsCount; i++ {
		tmp = strings.ReplaceAll(tmp, "{"+args[i]+"}", "")
		*mathContent = (*mathContent)[:idx] + tmp
	}
}
