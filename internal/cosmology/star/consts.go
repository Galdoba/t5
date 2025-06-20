package star

type SystemPosition string

var Primary = SystemPosition("P")
var PrimaryCompanion = SystemPosition("Pc")
var Close = SystemPosition("C")
var CloseCompanion = SystemPosition("Cc")
var Near = SystemPosition("N")
var NearCompanion = SystemPosition("Nc")
var Far = SystemPosition("F")
var FarCompanion = SystemPosition("Fc")
var Rogue = SystemPosition("R")

type SpectralType string

var SpectalType_O = SpectralType("O")
var SpectalType_B = SpectralType("B")
var SpectalType_A = SpectralType("A")
var SpectalType_F = SpectralType("F")
var SpectalType_G = SpectralType("G")
var SpectalType_K = SpectralType("K")
var SpectalType_M = SpectralType("M")
var SpectalType_BD = SpectralType("BD")
var SpectalType_Undefined = SpectralType("?")

type SizeClass string

var Size_Ia = SizeClass("Ia")
var Size_Ib = SizeClass("Ib")
var Size_II = SizeClass("II")
var Size_III = SizeClass("III")
var Size_IV = SizeClass("IV")
var Size_V = SizeClass("V")
var Size_VI = SizeClass("VI")
var Size_D = SizeClass("D")
var Size_NUL = SizeClass("")

type SubType string

var SubType_0 = SubType("0")
var SubType_1 = SubType("1")
var SubType_2 = SubType("2")
var SubType_3 = SubType("3")
var SubType_4 = SubType("4")
var SubType_5 = SubType("5")
var SubType_6 = SubType("6")
var SubType_7 = SubType("7")
var SubType_8 = SubType("8")
var SubType_9 = SubType("9")
var SubType_NUL = SubType("")

var GenerationMethod_T5basic = "T5 basic"
