package libvnc

/* $XConsortium: keysym.h,v 1.15 94/04/17 20:10:55 rws Exp $ */

const XK_VoidSymbol = 0xFFFFFF /* void symbol */

/*
 * TTY Functions, cleverly chosen to map to ascii, for convenience of
 * programming, but could have been arbitrary (at the cost of lookup
 * tables in client code.
 */

const XK_BackSpace = 0xFF08 /* back space, back char */
const XK_Tab = 0xFF09
const XK_Linefeed = 0xFF0A /* Linefeed, LF */
const XK_Clear = 0xFF0B
const XK_Return = 0xFF0D /* Return, enter */
const XK_Pause = 0xFF13  /* Pause, hold */
const XK_Scroll_Lock = 0xFF14
const XK_Sys_Req = 0xFF15
const XK_Escape = 0xFF1B
const XK_Delete = 0xFFFF /* Delete, rubout */

/* International & multi-key character composition */

const XK_Multi_key = 0xFF20 /* Multi-key character compose */
const XK_SingleCandidate = 0xFF3C
const XK_MultipleCandidate = 0xFF3D
const XK_PreviousCandidate = 0xFF3E

/* Japanese keyboard support */

const XK_Kanji = 0xFF21             /* Kanji, Kanji convert */
const XK_Muhenkan = 0xFF22          /* Cancel Conversion */
const XK_Henkan_Mode = 0xFF23       /* Start/Stop Conversion */
const XK_Henkan = 0xFF23            /* Alias for Henkan_Mode */
const XK_Romaji = 0xFF24            /* to Romaji */
const XK_Hiragana = 0xFF25          /* to Hiragana */
const XK_Katakana = 0xFF26          /* to Katakana */
const XK_Hiragana_Katakana = 0xFF27 /* Hiragana/Katakana toggle */
const XK_Zenkaku = 0xFF28           /* to Zenkaku */
const XK_Hankaku = 0xFF29           /* to Hankaku */
const XK_Zenkaku_Hankaku = 0xFF2A   /* Zenkaku/Hankaku toggle */
const XK_Touroku = 0xFF2B           /* Add to Dictionary */
const XK_Massyo = 0xFF2C            /* Delete from Dictionary */
const XK_Kana_Lock = 0xFF2D         /* Kana Lock */
const XK_Kana_Shift = 0xFF2E        /* Kana Shift */
const XK_Eisu_Shift = 0xFF2F        /* Alphanumeric Shift */
const XK_Eisu_toggle = 0xFF30       /* Alphanumeric toggle */
const XK_Zen_Koho = 0xFF3D          /* Multiple/All Candidate(s) */
const XK_Mae_Koho = 0xFF3E          /* Previous Candidate */

/*  = 0xFF31 through  = 0xFF3F are under XK_KOREAN */

/* Cursor control & motion */

const XK_Home = 0xFF50
const XK_Left = 0xFF51  /* Move left, left arrow */
const XK_Up = 0xFF52    /* Move up, up arrow */
const XK_Right = 0xFF53 /* Move right, right arrow */
const XK_Down = 0xFF54  /* Move down, down arrow */
const XK_Prior = 0xFF55 /* Prior, previous */
const XK_Page_Up = 0xFF55
const XK_Next = 0xFF56 /* Next */
const XK_Page_Down = 0xFF56
const XK_End = 0xFF57   /* EOL */
const XK_Begin = 0xFF58 /* BOL */

/* Misc Functions */

const XK_Select = 0xFF60 /* Select, mark */
const XK_Print = 0xFF61
const XK_Execute = 0xFF62 /* Execute, run, do */
const XK_Insert = 0xFF63  /* Insert, insert here */
const XK_Undo = 0xFF65    /* Undo, oops */
const XK_Redo = 0xFF66    /* redo, again */
const XK_Menu = 0xFF67
const XK_Find = 0xFF68   /* Find, search */
const XK_Cancel = 0xFF69 /* Cancel, stop, abort, exit */
const XK_Help = 0xFF6A   /* Help */
const XK_Break = 0xFF6B
const XK_Mode_switch = 0xFF7E   /* Character set switch */
const XK_script_switch = 0xFF7E /* Alias for mode_switch */
const XK_Num_Lock = 0xFF7F

/* Keypad Functions, keypad numbers cleverly chosen to map to ascii */

const XK_KP_Space = 0xFF80 /* space */
const XK_KP_Tab = 0xFF89
const XK_KP_Enter = 0xFF8D /* enter */
const XK_KP_F1 = 0xFF91    /* PF1, KP_A, ... */
const XK_KP_F2 = 0xFF92
const XK_KP_F3 = 0xFF93
const XK_KP_F4 = 0xFF94
const XK_KP_Home = 0xFF95
const XK_KP_Left = 0xFF96
const XK_KP_Up = 0xFF97
const XK_KP_Right = 0xFF98
const XK_KP_Down = 0xFF99
const XK_KP_Prior = 0xFF9A
const XK_KP_Page_Up = 0xFF9A
const XK_KP_Next = 0xFF9B
const XK_KP_Page_Down = 0xFF9B
const XK_KP_End = 0xFF9C
const XK_KP_Begin = 0xFF9D
const XK_KP_Insert = 0xFF9E
const XK_KP_Delete = 0xFF9F
const XK_KP_Equal = 0xFFBD /* equals */
const XK_KP_Multiply = 0xFFAA
const XK_KP_Add = 0xFFAB
const XK_KP_Separator = 0xFFAC /* separator, often comma */
const XK_KP_Subtract = 0xFFAD
const XK_KP_Decimal = 0xFFAE
const XK_KP_Divide = 0xFFAF

const XK_KP_0 = 0xFFB0
const XK_KP_1 = 0xFFB1
const XK_KP_2 = 0xFFB2
const XK_KP_3 = 0xFFB3
const XK_KP_4 = 0xFFB4
const XK_KP_5 = 0xFFB5
const XK_KP_6 = 0xFFB6
const XK_KP_7 = 0xFFB7
const XK_KP_8 = 0xFFB8
const XK_KP_9 = 0xFFB9

/*
 * Auxiliary Functions; note the duplicate definitions for left and right
 * function keys;  Sun keyboards and a few other manufactures have such
 * function key groups on the left and/or right sides of the keyboard.
 * We've not found a keyboard with more than 35 function keys total.
 */

const XK_F1 = 0xFFBE
const XK_F2 = 0xFFBF
const XK_F3 = 0xFFC0
const XK_F4 = 0xFFC1
const XK_F5 = 0xFFC2
const XK_F6 = 0xFFC3
const XK_F7 = 0xFFC4
const XK_F8 = 0xFFC5
const XK_F9 = 0xFFC6
const XK_F10 = 0xFFC7
const XK_F11 = 0xFFC8
const XK_L1 = 0xFFC8
const XK_F12 = 0xFFC9
const XK_L2 = 0xFFC9
const XK_F13 = 0xFFCA
const XK_L3 = 0xFFCA
const XK_F14 = 0xFFCB
const XK_L4 = 0xFFCB
const XK_F15 = 0xFFCC
const XK_L5 = 0xFFCC
const XK_F16 = 0xFFCD
const XK_L6 = 0xFFCD
const XK_F17 = 0xFFCE
const XK_L7 = 0xFFCE
const XK_F18 = 0xFFCF
const XK_L8 = 0xFFCF
const XK_F19 = 0xFFD0
const XK_L9 = 0xFFD0
const XK_F20 = 0xFFD1
const XK_L10 = 0xFFD1
const XK_F21 = 0xFFD2
const XK_R1 = 0xFFD2
const XK_F22 = 0xFFD3
const XK_R2 = 0xFFD3
const XK_F23 = 0xFFD4
const XK_R3 = 0xFFD4
const XK_F24 = 0xFFD5
const XK_R4 = 0xFFD5
const XK_F25 = 0xFFD6
const XK_R5 = 0xFFD6
const XK_F26 = 0xFFD7
const XK_R6 = 0xFFD7
const XK_F27 = 0xFFD8
const XK_R7 = 0xFFD8
const XK_F28 = 0xFFD9
const XK_R8 = 0xFFD9
const XK_F29 = 0xFFDA
const XK_R9 = 0xFFDA
const XK_F30 = 0xFFDB
const XK_R10 = 0xFFDB
const XK_F31 = 0xFFDC
const XK_R11 = 0xFFDC
const XK_F32 = 0xFFDD
const XK_R12 = 0xFFDD
const XK_F33 = 0xFFDE
const XK_R13 = 0xFFDE
const XK_F34 = 0xFFDF
const XK_R14 = 0xFFDF
const XK_F35 = 0xFFE0
const XK_R15 = 0xFFE0

/* Modifiers */

const XK_Shift_L = 0xFFE1    /* Left shift */
const XK_Shift_R = 0xFFE2    /* Right shift */
const XK_Control_L = 0xFFE3  /* Left control */
const XK_Control_R = 0xFFE4  /* Right control */
const XK_Caps_Lock = 0xFFE5  /* Caps lock */
const XK_Shift_Lock = 0xFFE6 /* Shift lock */

const XK_Meta_L = 0xFFE7  /* Left meta */
const XK_Meta_R = 0xFFE8  /* Right meta */
const XK_Alt_L = 0xFFE9   /* Left alt */
const XK_Alt_R = 0xFFEA   /* Right alt */
const XK_Super_L = 0xFFEB /* Left super */
const XK_Super_R = 0xFFEC /* Right super */
const XK_Hyper_L = 0xFFED /* Left hyper */
const XK_Hyper_R = 0xFFEE /* Right hyper */
/* XK_MISCELLANY */

/*
 * ISO 9995 Function and Modifier Keys
 * Byte 3 =  = 0xFE
 */

const XK_ISO_Lock = 0xFE01
const XK_ISO_Level2_Latch = 0xFE02
const XK_ISO_Level3_Shift = 0xFE03
const XK_ISO_Level3_Latch = 0xFE04
const XK_ISO_Level3_Lock = 0xFE05
const XK_ISO_Group_Shift = 0xFF7E /* Alias for mode_switch */
const XK_ISO_Group_Latch = 0xFE06
const XK_ISO_Group_Lock = 0xFE07
const XK_ISO_Next_Group = 0xFE08
const XK_ISO_Next_Group_Lock = 0xFE09
const XK_ISO_Prev_Group = 0xFE0A
const XK_ISO_Prev_Group_Lock = 0xFE0B
const XK_ISO_First_Group = 0xFE0C
const XK_ISO_First_Group_Lock = 0xFE0D
const XK_ISO_Last_Group = 0xFE0E
const XK_ISO_Last_Group_Lock = 0xFE0F

const XK_ISO_Left_Tab = 0xFE20
const XK_ISO_Move_Line_Up = 0xFE21
const XK_ISO_Move_Line_Down = 0xFE22
const XK_ISO_Partial_Line_Up = 0xFE23
const XK_ISO_Partial_Line_Down = 0xFE24
const XK_ISO_Partial_Space_Left = 0xFE25
const XK_ISO_Partial_Space_Right = 0xFE26
const XK_ISO_Set_Margin_Left = 0xFE27
const XK_ISO_Set_Margin_Right = 0xFE28
const XK_ISO_Release_Margin_Left = 0xFE29
const XK_ISO_Release_Margin_Right = 0xFE2A
const XK_ISO_Release_Both_Margins = 0xFE2B
const XK_ISO_Fast_Cursor_Left = 0xFE2C
const XK_ISO_Fast_Cursor_Right = 0xFE2D
const XK_ISO_Fast_Cursor_Up = 0xFE2E
const XK_ISO_Fast_Cursor_Down = 0xFE2F
const XK_ISO_Continuous_Underline = 0xFE30
const XK_ISO_Discontinuous_Underline = 0xFE31
const XK_ISO_Emphasize = 0xFE32
const XK_ISO_Center_Object = 0xFE33
const XK_ISO_Enter = 0xFE34

const XK_dead_grave = 0xFE50
const XK_dead_acute = 0xFE51
const XK_dead_circumflex = 0xFE52
const XK_dead_tilde = 0xFE53
const XK_dead_macron = 0xFE54
const XK_dead_breve = 0xFE55
const XK_dead_abovedot = 0xFE56
const XK_dead_diaeresis = 0xFE57
const XK_dead_abovering = 0xFE58
const XK_dead_doubleacute = 0xFE59
const XK_dead_caron = 0xFE5A
const XK_dead_cedilla = 0xFE5B
const XK_dead_ogonek = 0xFE5C
const XK_dead_iota = 0xFE5D
const XK_dead_voiced_sound = 0xFE5E
const XK_dead_semivoiced_sound = 0xFE5F
const XK_dead_belowdot = 0xFE60

const XK_First_Virtual_Screen = 0xFED0
const XK_Prev_Virtual_Screen = 0xFED1
const XK_Next_Virtual_Screen = 0xFED2
const XK_Last_Virtual_Screen = 0xFED4
const XK_Terminate_Server = 0xFED5

const XK_AccessX_Enable = 0xFE70
const XK_AccessX_Feedback_Enable = 0xFE71
const XK_RepeatKeys_Enable = 0xFE72
const XK_SlowKeys_Enable = 0xFE73
const XK_BounceKeys_Enable = 0xFE74
const XK_StickyKeys_Enable = 0xFE75
const XK_MouseKeys_Enable = 0xFE76
const XK_MouseKeys_Accel_Enable = 0xFE77
const XK_Overlay1_Enable = 0xFE78
const XK_Overlay2_Enable = 0xFE79
const XK_AudibleBell_Enable = 0xFE7A

const XK_Pointer_Left = 0xFEE0
const XK_Pointer_Right = 0xFEE1
const XK_Pointer_Up = 0xFEE2
const XK_Pointer_Down = 0xFEE3
const XK_Pointer_UpLeft = 0xFEE4
const XK_Pointer_UpRight = 0xFEE5
const XK_Pointer_DownLeft = 0xFEE6
const XK_Pointer_DownRight = 0xFEE7
const XK_Pointer_Button_Dflt = 0xFEE8
const XK_Pointer_Button1 = 0xFEE9
const XK_Pointer_Button2 = 0xFEEA
const XK_Pointer_Button3 = 0xFEEB
const XK_Pointer_Button4 = 0xFEEC
const XK_Pointer_Button5 = 0xFEED
const XK_Pointer_DblClick_Dflt = 0xFEEE
const XK_Pointer_DblClick1 = 0xFEEF
const XK_Pointer_DblClick2 = 0xFEF0
const XK_Pointer_DblClick3 = 0xFEF1
const XK_Pointer_DblClick4 = 0xFEF2
const XK_Pointer_DblClick5 = 0xFEF3
const XK_Pointer_Drag_Dflt = 0xFEF4
const XK_Pointer_Drag1 = 0xFEF5
const XK_Pointer_Drag2 = 0xFEF6
const XK_Pointer_Drag3 = 0xFEF7
const XK_Pointer_Drag4 = 0xFEF8
const XK_Pointer_Drag5 = 0xFEFD

const XK_Pointer_EnableKeys = 0xFEF9
const XK_Pointer_Accelerate = 0xFEFA
const XK_Pointer_DfltBtnNext = 0xFEFB
const XK_Pointer_DfltBtnPrev = 0xFEFC

/*
 * 3270 Terminal Keys
 * Byte 3 =  = 0xFD
 */
const XK_3270_Duplicate = 0xFD01
const XK_3270_FieldMark = 0xFD02
const XK_3270_Right2 = 0xFD03
const XK_3270_Left2 = 0xFD04
const XK_3270_BackTab = 0xFD05
const XK_3270_EraseEOF = 0xFD06
const XK_3270_EraseInput = 0xFD07
const XK_3270_Reset = 0xFD08
const XK_3270_Quit = 0xFD09
const XK_3270_PA1 = 0xFD0A
const XK_3270_PA2 = 0xFD0B
const XK_3270_PA3 = 0xFD0C
const XK_3270_Test = 0xFD0D
const XK_3270_Attn = 0xFD0E
const XK_3270_CursorBlink = 0xFD0F
const XK_3270_AltCursor = 0xFD10
const XK_3270_KeyClick = 0xFD11
const XK_3270_Jump = 0xFD12
const XK_3270_Ident = 0xFD13
const XK_3270_Rule = 0xFD14
const XK_3270_Copy = 0xFD15
const XK_3270_Play = 0xFD16
const XK_3270_Setup = 0xFD17
const XK_3270_Record = 0xFD18
const XK_3270_ChangeScreen = 0xFD19
const XK_3270_DeleteWord = 0xFD1A
const XK_3270_ExSelect = 0xFD1B
const XK_3270_CursorSelect = 0xFD1C
const XK_3270_PrintScreen = 0xFD1D
const XK_3270_Enter = 0xFD1E

/*
 *  Latin 1
 *  Byte 3 = 0
 */

const XK_space = 0x020
const XK_exclam = 0x021
const XK_quotedbl = 0x022
const XK_numbersign = 0x023
const XK_dollar = 0x024
const XK_percent = 0x025
const XK_ampersand = 0x026
const XK_apostrophe = 0x027
const XK_quoteright = 0x027 /* deprecated */
const XK_parenleft = 0x028
const XK_parenright = 0x029
const XK_asterisk = 0x02a
const XK_plus = 0x02b
const XK_comma = 0x02c
const XK_minus = 0x02d
const XK_period = 0x02e
const XK_slash = 0x02f
const XK_0 = 0x030
const XK_1 = 0x031
const XK_2 = 0x032
const XK_3 = 0x033
const XK_4 = 0x034
const XK_5 = 0x035
const XK_6 = 0x036
const XK_7 = 0x037
const XK_8 = 0x038
const XK_9 = 0x039
const XK_colon = 0x03a
const XK_semicolon = 0x03b
const XK_less = 0x03c
const XK_equal = 0x03d
const XK_greater = 0x03e
const XK_question = 0x03f
const XK_at = 0x040
const XK_A = 0x041
const XK_B = 0x042
const XK_C = 0x043
const XK_D = 0x044
const XK_E = 0x045
const XK_F = 0x046
const XK_G = 0x047
const XK_H = 0x048
const XK_I = 0x049
const XK_J = 0x04a
const XK_K = 0x04b
const XK_L = 0x04c
const XK_M = 0x04d
const XK_N = 0x04e
const XK_O = 0x04f
const XK_P = 0x050
const XK_Q = 0x051
const XK_R = 0x052
const XK_S = 0x053
const XK_T = 0x054
const XK_U = 0x055
const XK_V = 0x056
const XK_W = 0x057
const XK_X = 0x058
const XK_Y = 0x059
const XK_Z = 0x05a
const XK_bracketleft = 0x05b
const XK_backslash = 0x05c
const XK_bracketright = 0x05d
const XK_asciicircum = 0x05e
const XK_underscore = 0x05f
const XK_grave = 0x060
const XK_quoteleft = 0x060 /* deprecated */
const XK_a = 0x061
const XK_b = 0x062
const XK_c = 0x063
const XK_d = 0x064
const XK_e = 0x065
const XK_f = 0x066
const XK_g = 0x067
const XK_h = 0x068
const XK_i = 0x069
const XK_j = 0x06a
const XK_k = 0x06b
const XK_l = 0x06c
const XK_m = 0x06d
const XK_n = 0x06e
const XK_o = 0x06f
const XK_p = 0x070
const XK_q = 0x071
const XK_r = 0x072
const XK_s = 0x073
const XK_t = 0x074
const XK_u = 0x075
const XK_v = 0x076
const XK_w = 0x077
const XK_x = 0x078
const XK_y = 0x079
const XK_z = 0x07a
const XK_braceleft = 0x07b
const XK_bar = 0x07c
const XK_braceright = 0x07d
const XK_asciitilde = 0x07e

const XK_nobreakspace = 0x0a0
const XK_exclamdown = 0x0a1
const XK_cent = 0x0a2
const XK_sterling = 0x0a3
const XK_currency = 0x0a4
const XK_yen = 0x0a5
const XK_brokenbar = 0x0a6
const XK_section = 0x0a7
const XK_diaeresis = 0x0a8
const XK_copyright = 0x0a9
const XK_ordfeminine = 0x0aa
const XK_guillemotleft = 0x0ab /* left angle quotation mark */
const XK_notsign = 0x0ac
const XK_hyphen = 0x0ad
const XK_registered = 0x0ae
const XK_macron = 0x0af
const XK_degree = 0x0b0
const XK_plusminus = 0x0b1
const XK_twosuperior = 0x0b2
const XK_threesuperior = 0x0b3
const XK_acute = 0x0b4
const XK_mu = 0x0b5
const XK_paragraph = 0x0b6
const XK_periodcentered = 0x0b7
const XK_cedilla = 0x0b8
const XK_onesuperior = 0x0b9
const XK_masculine = 0x0ba
const XK_guillemotright = 0x0bb /* right angle quotation mark */
const XK_onequarter = 0x0bc
const XK_onehalf = 0x0bd
const XK_threequarters = 0x0be
const XK_questiondown = 0x0bf
const XK_Agrave = 0x0c0
const XK_Aacute = 0x0c1
const XK_Acircumflex = 0x0c2
const XK_Atilde = 0x0c3
const XK_Adiaeresis = 0x0c4
const XK_Aring = 0x0c5
const XK_AE = 0x0c6
const XK_Ccedilla = 0x0c7
const XK_Egrave = 0x0c8
const XK_Eacute = 0x0c9
const XK_Ecircumflex = 0x0ca
const XK_Ediaeresis = 0x0cb
const XK_Igrave = 0x0cc
const XK_Iacute = 0x0cd
const XK_Icircumflex = 0x0ce
const XK_Idiaeresis = 0x0cf
const XK_ETH = 0x0d0
const XK_Eth = 0x0d0 /* deprecated */
const XK_Ntilde = 0x0d1
const XK_Ograve = 0x0d2
const XK_Oacute = 0x0d3
const XK_Ocircumflex = 0x0d4
const XK_Otilde = 0x0d5
const XK_Odiaeresis = 0x0d6
const XK_multiply = 0x0d7
const XK_Ooblique = 0x0d8
const XK_Ugrave = 0x0d9
const XK_Uacute = 0x0da
const XK_Ucircumflex = 0x0db
const XK_Udiaeresis = 0x0dc
const XK_Yacute = 0x0dd
const XK_THORN = 0x0de
const XK_Thorn = 0x0de /* deprecated */
const XK_ssharp = 0x0df
const XK_agrave = 0x0e0
const XK_aacute = 0x0e1
const XK_acircumflex = 0x0e2
const XK_atilde = 0x0e3
const XK_adiaeresis = 0x0e4
const XK_aring = 0x0e5
const XK_ae = 0x0e6
const XK_ccedilla = 0x0e7
const XK_egrave = 0x0e8
const XK_eacute = 0x0e9
const XK_ecircumflex = 0x0ea
const XK_ediaeresis = 0x0eb
const XK_igrave = 0x0ec
const XK_iacute = 0x0ed
const XK_icircumflex = 0x0ee
const XK_idiaeresis = 0x0ef
const XK_eth = 0x0f0
const XK_ntilde = 0x0f1
const XK_ograve = 0x0f2
const XK_oacute = 0x0f3
const XK_ocircumflex = 0x0f4
const XK_otilde = 0x0f5
const XK_odiaeresis = 0x0f6
const XK_division = 0x0f7
const XK_oslash = 0x0f8
const XK_ugrave = 0x0f9
const XK_uacute = 0x0fa
const XK_ucircumflex = 0x0fb
const XK_udiaeresis = 0x0fc
const XK_yacute = 0x0fd
const XK_thorn = 0x0fe
const XK_ydiaeresis = 0x0ff

/* XK_LATIN1 */

/*
 *   Latin 2
 *   Byte 3 = 1
 */

const XK_Aogonek = 0x1a1
const XK_breve = 0x1a2
const XK_Lstroke = 0x1a3
const XK_Lcaron = 0x1a5
const XK_Sacute = 0x1a6
const XK_Scaron = 0x1a9
const XK_Scedilla = 0x1aa
const XK_Tcaron = 0x1ab
const XK_Zacute = 0x1ac
const XK_Zcaron = 0x1ae
const XK_Zabovedot = 0x1af
const XK_aogonek = 0x1b1
const XK_ogonek = 0x1b2
const XK_lstroke = 0x1b3
const XK_lcaron = 0x1b5
const XK_sacute = 0x1b6
const XK_caron = 0x1b7
const XK_scaron = 0x1b9
const XK_scedilla = 0x1ba
const XK_tcaron = 0x1bb
const XK_zacute = 0x1bc
const XK_doubleacute = 0x1bd
const XK_zcaron = 0x1be
const XK_zabovedot = 0x1bf
const XK_Racute = 0x1c0
const XK_Abreve = 0x1c3
const XK_Lacute = 0x1c5
const XK_Cacute = 0x1c6
const XK_Ccaron = 0x1c8
const XK_Eogonek = 0x1ca
const XK_Ecaron = 0x1cc
const XK_Dcaron = 0x1cf
const XK_Dstroke = 0x1d0
const XK_Nacute = 0x1d1
const XK_Ncaron = 0x1d2
const XK_Odoubleacute = 0x1d5
const XK_Rcaron = 0x1d8
const XK_Uring = 0x1d9
const XK_Udoubleacute = 0x1db
const XK_Tcedilla = 0x1de
const XK_racute = 0x1e0
const XK_abreve = 0x1e3
const XK_lacute = 0x1e5
const XK_cacute = 0x1e6
const XK_ccaron = 0x1e8
const XK_eogonek = 0x1ea
const XK_ecaron = 0x1ec
const XK_dcaron = 0x1ef
const XK_dstroke = 0x1f0
const XK_nacute = 0x1f1
const XK_ncaron = 0x1f2
const XK_odoubleacute = 0x1f5
const XK_udoubleacute = 0x1fb
const XK_rcaron = 0x1f8
const XK_uring = 0x1f9
const XK_tcedilla = 0x1fe
const XK_abovedot = 0x1ff

/* XK_LATIN2 */

/*
 *   Latin 3
 *   Byte 3 = 2
 */

const XK_Hstroke = 0x2a1
const XK_Hcircumflex = 0x2a6
const XK_Iabovedot = 0x2a9
const XK_Gbreve = 0x2ab
const XK_Jcircumflex = 0x2ac
const XK_hstroke = 0x2b1
const XK_hcircumflex = 0x2b6
const XK_idotless = 0x2b9
const XK_gbreve = 0x2bb
const XK_jcircumflex = 0x2bc
const XK_Cabovedot = 0x2c5
const XK_Ccircumflex = 0x2c6
const XK_Gabovedot = 0x2d5
const XK_Gcircumflex = 0x2d8
const XK_Ubreve = 0x2dd
const XK_Scircumflex = 0x2de
const XK_cabovedot = 0x2e5
const XK_ccircumflex = 0x2e6
const XK_gabovedot = 0x2f5
const XK_gcircumflex = 0x2f8
const XK_ubreve = 0x2fd
const XK_scircumflex = 0x2fe

/* XK_LATIN3 */

/*
 *   Latin 4
 *   Byte 3 = 3
 */

const XK_kra = 0x3a2
const XK_kappa = 0x3a2 /* deprecated */
const XK_Rcedilla = 0x3a3
const XK_Itilde = 0x3a5
const XK_Lcedilla = 0x3a6
const XK_Emacron = 0x3aa
const XK_Gcedilla = 0x3ab
const XK_Tslash = 0x3ac
const XK_rcedilla = 0x3b3
const XK_itilde = 0x3b5
const XK_lcedilla = 0x3b6
const XK_emacron = 0x3ba
const XK_gcedilla = 0x3bb
const XK_tslash = 0x3bc
const XK_ENG = 0x3bd
const XK_eng = 0x3bf
const XK_Amacron = 0x3c0
const XK_Iogonek = 0x3c7
const XK_Eabovedot = 0x3cc
const XK_Imacron = 0x3cf
const XK_Ncedilla = 0x3d1
const XK_Omacron = 0x3d2
const XK_Kcedilla = 0x3d3
const XK_Uogonek = 0x3d9
const XK_Utilde = 0x3dd
const XK_Umacron = 0x3de
const XK_amacron = 0x3e0
const XK_iogonek = 0x3e7
const XK_eabovedot = 0x3ec
const XK_imacron = 0x3ef
const XK_ncedilla = 0x3f1
const XK_omacron = 0x3f2
const XK_kcedilla = 0x3f3
const XK_uogonek = 0x3f9
const XK_utilde = 0x3fd
const XK_umacron = 0x3fe

/* XK_LATIN4 */

/*
 * Katakana
 * Byte 3 = 4
 */

const XK_overline = 0x47e
const XK_kana_fullstop = 0x4a1
const XK_kana_openingbracket = 0x4a2
const XK_kana_closingbracket = 0x4a3
const XK_kana_comma = 0x4a4
const XK_kana_conjunctive = 0x4a5
const XK_kana_middledot = 0x4a5 /* deprecated */
const XK_kana_WO = 0x4a6
const XK_kana_a = 0x4a7
const XK_kana_i = 0x4a8
const XK_kana_u = 0x4a9
const XK_kana_e = 0x4aa
const XK_kana_o = 0x4ab
const XK_kana_ya = 0x4ac
const XK_kana_yu = 0x4ad
const XK_kana_yo = 0x4ae
const XK_kana_tsu = 0x4af
const XK_kana_tu = 0x4af /* deprecated */
const XK_prolongedsound = 0x4b0
const XK_kana_A = 0x4b1
const XK_kana_I = 0x4b2
const XK_kana_U = 0x4b3
const XK_kana_E = 0x4b4
const XK_kana_O = 0x4b5
const XK_kana_KA = 0x4b6
const XK_kana_KI = 0x4b7
const XK_kana_KU = 0x4b8
const XK_kana_KE = 0x4b9
const XK_kana_KO = 0x4ba
const XK_kana_SA = 0x4bb
const XK_kana_SHI = 0x4bc
const XK_kana_SU = 0x4bd
const XK_kana_SE = 0x4be
const XK_kana_SO = 0x4bf
const XK_kana_TA = 0x4c0
const XK_kana_CHI = 0x4c1
const XK_kana_TI = 0x4c1 /* deprecated */
const XK_kana_TSU = 0x4c2
const XK_kana_TU = 0x4c2 /* deprecated */
const XK_kana_TE = 0x4c3
const XK_kana_TO = 0x4c4
const XK_kana_NA = 0x4c5
const XK_kana_NI = 0x4c6
const XK_kana_NU = 0x4c7
const XK_kana_NE = 0x4c8
const XK_kana_NO = 0x4c9
const XK_kana_HA = 0x4ca
const XK_kana_HI = 0x4cb
const XK_kana_FU = 0x4cc
const XK_kana_HU = 0x4cc /* deprecated */
const XK_kana_HE = 0x4cd
const XK_kana_HO = 0x4ce
const XK_kana_MA = 0x4cf
const XK_kana_MI = 0x4d0
const XK_kana_MU = 0x4d1
const XK_kana_ME = 0x4d2
const XK_kana_MO = 0x4d3
const XK_kana_YA = 0x4d4
const XK_kana_YU = 0x4d5
const XK_kana_YO = 0x4d6
const XK_kana_RA = 0x4d7
const XK_kana_RI = 0x4d8
const XK_kana_RU = 0x4d9
const XK_kana_RE = 0x4da
const XK_kana_RO = 0x4db
const XK_kana_WA = 0x4dc
const XK_kana_N = 0x4dd
const XK_voicedsound = 0x4de
const XK_semivoicedsound = 0x4df
const XK_kana_switch = 0xFF7E /* Alias for mode_switch */
/* XK_KATAKANA */

/*
 *  Arabic
 *  Byte 3 = 5
 */

const XK_Arabic_comma = 0x5ac
const XK_Arabic_semicolon = 0x5bb
const XK_Arabic_question_mark = 0x5bf
const XK_Arabic_hamza = 0x5c1
const XK_Arabic_maddaonalef = 0x5c2
const XK_Arabic_hamzaonalef = 0x5c3
const XK_Arabic_hamzaonwaw = 0x5c4
const XK_Arabic_hamzaunderalef = 0x5c5
const XK_Arabic_hamzaonyeh = 0x5c6
const XK_Arabic_alef = 0x5c7
const XK_Arabic_beh = 0x5c8
const XK_Arabic_tehmarbuta = 0x5c9
const XK_Arabic_teh = 0x5ca
const XK_Arabic_theh = 0x5cb
const XK_Arabic_jeem = 0x5cc
const XK_Arabic_hah = 0x5cd
const XK_Arabic_khah = 0x5ce
const XK_Arabic_dal = 0x5cf
const XK_Arabic_thal = 0x5d0
const XK_Arabic_ra = 0x5d1
const XK_Arabic_zain = 0x5d2
const XK_Arabic_seen = 0x5d3
const XK_Arabic_sheen = 0x5d4
const XK_Arabic_sad = 0x5d5
const XK_Arabic_dad = 0x5d6
const XK_Arabic_tah = 0x5d7
const XK_Arabic_zah = 0x5d8
const XK_Arabic_ain = 0x5d9
const XK_Arabic_ghain = 0x5da
const XK_Arabic_tatweel = 0x5e0
const XK_Arabic_feh = 0x5e1
const XK_Arabic_qaf = 0x5e2
const XK_Arabic_kaf = 0x5e3
const XK_Arabic_lam = 0x5e4
const XK_Arabic_meem = 0x5e5
const XK_Arabic_noon = 0x5e6
const XK_Arabic_ha = 0x5e7
const XK_Arabic_heh = 0x5e7 /* deprecated */
const XK_Arabic_waw = 0x5e8
const XK_Arabic_alefmaksura = 0x5e9
const XK_Arabic_yeh = 0x5ea
const XK_Arabic_fathatan = 0x5eb
const XK_Arabic_dammatan = 0x5ec
const XK_Arabic_kasratan = 0x5ed
const XK_Arabic_fatha = 0x5ee
const XK_Arabic_damma = 0x5ef
const XK_Arabic_kasra = 0x5f0
const XK_Arabic_shadda = 0x5f1
const XK_Arabic_sukun = 0x5f2
const XK_Arabic_switch = 0xFF7E /* Alias for mode_switch */
/* XK_ARABIC */

/*
 * Cyrillic
 * Byte 3 = 6
 */
const XK_Serbian_dje = 0x6a1
const XK_Macedonia_gje = 0x6a2
const XK_Cyrillic_io = 0x6a3
const XK_Ukrainian_ie = 0x6a4
const XK_Ukranian_je = 0x6a4 /* deprecated */
const XK_Macedonia_dse = 0x6a5
const XK_Ukrainian_i = 0x6a6
const XK_Ukranian_i = 0x6a6 /* deprecated */
const XK_Ukrainian_yi = 0x6a7
const XK_Ukranian_yi = 0x6a7 /* deprecated */
const XK_Cyrillic_je = 0x6a8
const XK_Serbian_je = 0x6a8 /* deprecated */
const XK_Cyrillic_lje = 0x6a9
const XK_Serbian_lje = 0x6a9 /* deprecated */
const XK_Cyrillic_nje = 0x6aa
const XK_Serbian_nje = 0x6aa /* deprecated */
const XK_Serbian_tshe = 0x6ab
const XK_Macedonia_kje = 0x6ac
const XK_Byelorussian_shortu = 0x6ae
const XK_Cyrillic_dzhe = 0x6af
const XK_Serbian_dze = 0x6af /* deprecated */
const XK_numerosign = 0x6b0
const XK_Serbian_DJE = 0x6b1
const XK_Macedonia_GJE = 0x6b2
const XK_Cyrillic_IO = 0x6b3
const XK_Ukrainian_IE = 0x6b4
const XK_Ukranian_JE = 0x6b4 /* deprecated */
const XK_Macedonia_DSE = 0x6b5
const XK_Ukrainian_I = 0x6b6
const XK_Ukranian_I = 0x6b6 /* deprecated */
const XK_Ukrainian_YI = 0x6b7
const XK_Ukranian_YI = 0x6b7 /* deprecated */
const XK_Cyrillic_JE = 0x6b8
const XK_Serbian_JE = 0x6b8 /* deprecated */
const XK_Cyrillic_LJE = 0x6b9
const XK_Serbian_LJE = 0x6b9 /* deprecated */
const XK_Cyrillic_NJE = 0x6ba
const XK_Serbian_NJE = 0x6ba /* deprecated */
const XK_Serbian_TSHE = 0x6bb
const XK_Macedonia_KJE = 0x6bc
const XK_Byelorussian_SHORTU = 0x6be
const XK_Cyrillic_DZHE = 0x6bf
const XK_Serbian_DZE = 0x6bf /* deprecated */
const XK_Cyrillic_yu = 0x6c0
const XK_Cyrillic_a = 0x6c1
const XK_Cyrillic_be = 0x6c2
const XK_Cyrillic_tse = 0x6c3
const XK_Cyrillic_de = 0x6c4
const XK_Cyrillic_ie = 0x6c5
const XK_Cyrillic_ef = 0x6c6
const XK_Cyrillic_ghe = 0x6c7
const XK_Cyrillic_ha = 0x6c8
const XK_Cyrillic_i = 0x6c9
const XK_Cyrillic_shorti = 0x6ca
const XK_Cyrillic_ka = 0x6cb
const XK_Cyrillic_el = 0x6cc
const XK_Cyrillic_em = 0x6cd
const XK_Cyrillic_en = 0x6ce
const XK_Cyrillic_o = 0x6cf
const XK_Cyrillic_pe = 0x6d0
const XK_Cyrillic_ya = 0x6d1
const XK_Cyrillic_er = 0x6d2
const XK_Cyrillic_es = 0x6d3
const XK_Cyrillic_te = 0x6d4
const XK_Cyrillic_u = 0x6d5
const XK_Cyrillic_zhe = 0x6d6
const XK_Cyrillic_ve = 0x6d7
const XK_Cyrillic_softsign = 0x6d8
const XK_Cyrillic_yeru = 0x6d9
const XK_Cyrillic_ze = 0x6da
const XK_Cyrillic_sha = 0x6db
const XK_Cyrillic_e = 0x6dc
const XK_Cyrillic_shcha = 0x6dd
const XK_Cyrillic_che = 0x6de
const XK_Cyrillic_hardsign = 0x6df
const XK_Cyrillic_YU = 0x6e0
const XK_Cyrillic_A = 0x6e1
const XK_Cyrillic_BE = 0x6e2
const XK_Cyrillic_TSE = 0x6e3
const XK_Cyrillic_DE = 0x6e4
const XK_Cyrillic_IE = 0x6e5
const XK_Cyrillic_EF = 0x6e6
const XK_Cyrillic_GHE = 0x6e7
const XK_Cyrillic_HA = 0x6e8
const XK_Cyrillic_I = 0x6e9
const XK_Cyrillic_SHORTI = 0x6ea
const XK_Cyrillic_KA = 0x6eb
const XK_Cyrillic_EL = 0x6ec
const XK_Cyrillic_EM = 0x6ed
const XK_Cyrillic_EN = 0x6ee
const XK_Cyrillic_O = 0x6ef
const XK_Cyrillic_PE = 0x6f0
const XK_Cyrillic_YA = 0x6f1
const XK_Cyrillic_ER = 0x6f2
const XK_Cyrillic_ES = 0x6f3
const XK_Cyrillic_TE = 0x6f4
const XK_Cyrillic_U = 0x6f5
const XK_Cyrillic_ZHE = 0x6f6
const XK_Cyrillic_VE = 0x6f7
const XK_Cyrillic_SOFTSIGN = 0x6f8
const XK_Cyrillic_YERU = 0x6f9
const XK_Cyrillic_ZE = 0x6fa
const XK_Cyrillic_SHA = 0x6fb
const XK_Cyrillic_E = 0x6fc
const XK_Cyrillic_SHCHA = 0x6fd
const XK_Cyrillic_CHE = 0x6fe
const XK_Cyrillic_HARDSIGN = 0x6ff

/* XK_CYRILLIC */

/*
 * Greek
 * Byte 3 = 7
 */

const XK_Greek_ALPHAaccent = 0x7a1
const XK_Greek_EPSILONaccent = 0x7a2
const XK_Greek_ETAaccent = 0x7a3
const XK_Greek_IOTAaccent = 0x7a4
const XK_Greek_IOTAdieresis = 0x7a5
const XK_Greek_OMICRONaccent = 0x7a7
const XK_Greek_UPSILONaccent = 0x7a8
const XK_Greek_UPSILONdieresis = 0x7a9
const XK_Greek_OMEGAaccent = 0x7ab
const XK_Greek_accentdieresis = 0x7ae
const XK_Greek_horizbar = 0x7af
const XK_Greek_alphaaccent = 0x7b1
const XK_Greek_epsilonaccent = 0x7b2
const XK_Greek_etaaccent = 0x7b3
const XK_Greek_iotaaccent = 0x7b4
const XK_Greek_iotadieresis = 0x7b5
const XK_Greek_iotaaccentdieresis = 0x7b6
const XK_Greek_omicronaccent = 0x7b7
const XK_Greek_upsilonaccent = 0x7b8
const XK_Greek_upsilondieresis = 0x7b9
const XK_Greek_upsilonaccentdieresis = 0x7ba
const XK_Greek_omegaaccent = 0x7bb
const XK_Greek_ALPHA = 0x7c1
const XK_Greek_BETA = 0x7c2
const XK_Greek_GAMMA = 0x7c3
const XK_Greek_DELTA = 0x7c4
const XK_Greek_EPSILON = 0x7c5
const XK_Greek_ZETA = 0x7c6
const XK_Greek_ETA = 0x7c7
const XK_Greek_THETA = 0x7c8
const XK_Greek_IOTA = 0x7c9
const XK_Greek_KAPPA = 0x7ca
const XK_Greek_LAMDA = 0x7cb
const XK_Greek_LAMBDA = 0x7cb
const XK_Greek_MU = 0x7cc
const XK_Greek_NU = 0x7cd
const XK_Greek_XI = 0x7ce
const XK_Greek_OMICRON = 0x7cf
const XK_Greek_PI = 0x7d0
const XK_Greek_RHO = 0x7d1
const XK_Greek_SIGMA = 0x7d2
const XK_Greek_TAU = 0x7d4
const XK_Greek_UPSILON = 0x7d5
const XK_Greek_PHI = 0x7d6
const XK_Greek_CHI = 0x7d7
const XK_Greek_PSI = 0x7d8
const XK_Greek_OMEGA = 0x7d9
const XK_Greek_alpha = 0x7e1
const XK_Greek_beta = 0x7e2
const XK_Greek_gamma = 0x7e3
const XK_Greek_delta = 0x7e4
const XK_Greek_epsilon = 0x7e5
const XK_Greek_zeta = 0x7e6
const XK_Greek_eta = 0x7e7
const XK_Greek_theta = 0x7e8
const XK_Greek_iota = 0x7e9
const XK_Greek_kappa = 0x7ea
const XK_Greek_lamda = 0x7eb
const XK_Greek_lambda = 0x7eb
const XK_Greek_mu = 0x7ec
const XK_Greek_nu = 0x7ed
const XK_Greek_xi = 0x7ee
const XK_Greek_omicron = 0x7ef
const XK_Greek_pi = 0x7f0
const XK_Greek_rho = 0x7f1
const XK_Greek_sigma = 0x7f2
const XK_Greek_finalsmallsigma = 0x7f3
const XK_Greek_tau = 0x7f4
const XK_Greek_upsilon = 0x7f5
const XK_Greek_phi = 0x7f6
const XK_Greek_chi = 0x7f7
const XK_Greek_psi = 0x7f8
const XK_Greek_omega = 0x7f9
const XK_Greek_switch = 0xFF7E /* Alias for mode_switch */
/* XK_GREEK */

/*
 * Technical
 * Byte 3 = 8
 */

const XK_leftradical = 0x8a1
const XK_topleftradical = 0x8a2
const XK_horizconnector = 0x8a3
const XK_topintegral = 0x8a4
const XK_botintegral = 0x8a5
const XK_vertconnector = 0x8a6
const XK_topleftsqbracket = 0x8a7
const XK_botleftsqbracket = 0x8a8
const XK_toprightsqbracket = 0x8a9
const XK_botrightsqbracket = 0x8aa
const XK_topleftparens = 0x8ab
const XK_botleftparens = 0x8ac
const XK_toprightparens = 0x8ad
const XK_botrightparens = 0x8ae
const XK_leftmiddlecurlybrace = 0x8af
const XK_rightmiddlecurlybrace = 0x8b0
const XK_topleftsummation = 0x8b1
const XK_botleftsummation = 0x8b2
const XK_topvertsummationconnector = 0x8b3
const XK_botvertsummationconnector = 0x8b4
const XK_toprightsummation = 0x8b5
const XK_botrightsummation = 0x8b6
const XK_rightmiddlesummation = 0x8b7
const XK_lessthanequal = 0x8bc
const XK_notequal = 0x8bd
const XK_greaterthanequal = 0x8be
const XK_integral = 0x8bf
const XK_therefore = 0x8c0
const XK_variation = 0x8c1
const XK_infinity = 0x8c2
const XK_nabla = 0x8c5
const XK_approximate = 0x8c8
const XK_similarequal = 0x8c9
const XK_ifonlyif = 0x8cd
const XK_implies = 0x8ce
const XK_identical = 0x8cf
const XK_radical = 0x8d6
const XK_includedin = 0x8da
const XK_includes = 0x8db
const XK_intersection = 0x8dc
const XK_union = 0x8dd
const XK_logicaland = 0x8de
const XK_logicalor = 0x8df
const XK_partialderivative = 0x8ef
const XK_function = 0x8f6
const XK_leftarrow = 0x8fb
const XK_uparrow = 0x8fc
const XK_rightarrow = 0x8fd
const XK_downarrow = 0x8fe

/* XK_TECHNICAL */

/*
 *  Special
 *  Byte 3 = 9
 */

const XK_blank = 0x9df
const XK_soliddiamond = 0x9e0
const XK_checkerboard = 0x9e1
const XK_ht = 0x9e2
const XK_ff = 0x9e3
const XK_cr = 0x9e4
const XK_lf = 0x9e5
const XK_nl = 0x9e8
const XK_vt = 0x9e9
const XK_lowrightcorner = 0x9ea
const XK_uprightcorner = 0x9eb
const XK_upleftcorner = 0x9ec
const XK_lowleftcorner = 0x9ed
const XK_crossinglines = 0x9ee
const XK_horizlinescan1 = 0x9ef
const XK_horizlinescan3 = 0x9f0
const XK_horizlinescan5 = 0x9f1
const XK_horizlinescan7 = 0x9f2
const XK_horizlinescan9 = 0x9f3
const XK_leftt = 0x9f4
const XK_rightt = 0x9f5
const XK_bott = 0x9f6
const XK_topt = 0x9f7
const XK_vertbar = 0x9f8

/* XK_SPECIAL */

/*
 *  Publishing
 *  Byte 3 = a
 */

const XK_emspace = 0xaa1
const XK_enspace = 0xaa2
const XK_em3space = 0xaa3
const XK_em4space = 0xaa4
const XK_digitspace = 0xaa5
const XK_punctspace = 0xaa6
const XK_thinspace = 0xaa7
const XK_hairspace = 0xaa8
const XK_emdash = 0xaa9
const XK_endash = 0xaaa
const XK_signifblank = 0xaac
const XK_ellipsis = 0xaae
const XK_doubbaselinedot = 0xaaf
const XK_onethird = 0xab0
const XK_twothirds = 0xab1
const XK_onefifth = 0xab2
const XK_twofifths = 0xab3
const XK_threefifths = 0xab4
const XK_fourfifths = 0xab5
const XK_onesixth = 0xab6
const XK_fivesixths = 0xab7
const XK_careof = 0xab8
const XK_figdash = 0xabb
const XK_leftanglebracket = 0xabc
const XK_decimalpoint = 0xabd
const XK_rightanglebracket = 0xabe
const XK_marker = 0xabf
const XK_oneeighth = 0xac3
const XK_threeeighths = 0xac4
const XK_fiveeighths = 0xac5
const XK_seveneighths = 0xac6
const XK_trademark = 0xac9
const XK_signaturemark = 0xaca
const XK_trademarkincircle = 0xacb
const XK_leftopentriangle = 0xacc
const XK_rightopentriangle = 0xacd
const XK_emopencircle = 0xace
const XK_emopenrectangle = 0xacf
const XK_leftsinglequotemark = 0xad0
const XK_rightsinglequotemark = 0xad1
const XK_leftdoublequotemark = 0xad2
const XK_rightdoublequotemark = 0xad3
const XK_prescription = 0xad4
const XK_minutes = 0xad6
const XK_seconds = 0xad7
const XK_latincross = 0xad9
const XK_hexagram = 0xada
const XK_filledrectbullet = 0xadb
const XK_filledlefttribullet = 0xadc
const XK_filledrighttribullet = 0xadd
const XK_emfilledcircle = 0xade
const XK_emfilledrect = 0xadf
const XK_enopencircbullet = 0xae0
const XK_enopensquarebullet = 0xae1
const XK_openrectbullet = 0xae2
const XK_opentribulletup = 0xae3
const XK_opentribulletdown = 0xae4
const XK_openstar = 0xae5
const XK_enfilledcircbullet = 0xae6
const XK_enfilledsqbullet = 0xae7
const XK_filledtribulletup = 0xae8
const XK_filledtribulletdown = 0xae9
const XK_leftpointer = 0xaea
const XK_rightpointer = 0xaeb
const XK_club = 0xaec
const XK_diamond = 0xaed
const XK_heart = 0xaee
const XK_maltesecross = 0xaf0
const XK_dagger = 0xaf1
const XK_doubledagger = 0xaf2
const XK_checkmark = 0xaf3
const XK_ballotcross = 0xaf4
const XK_musicalsharp = 0xaf5
const XK_musicalflat = 0xaf6
const XK_malesymbol = 0xaf7
const XK_femalesymbol = 0xaf8
const XK_telephone = 0xaf9
const XK_telephonerecorder = 0xafa
const XK_phonographcopyright = 0xafb
const XK_caret = 0xafc
const XK_singlelowquotemark = 0xafd
const XK_doublelowquotemark = 0xafe
const XK_cursor = 0xaff

/* XK_PUBLISHING */

/*
 *  APL
 *  Byte 3 = b
 */

const XK_leftcaret = 0xba3
const XK_rightcaret = 0xba6
const XK_downcaret = 0xba8
const XK_upcaret = 0xba9
const XK_overbar = 0xbc0
const XK_downtack = 0xbc2
const XK_upshoe = 0xbc3
const XK_downstile = 0xbc4
const XK_underbar = 0xbc6
const XK_jot = 0xbca
const XK_quad = 0xbcc
const XK_uptack = 0xbce
const XK_circle = 0xbcf
const XK_upstile = 0xbd3
const XK_downshoe = 0xbd6
const XK_rightshoe = 0xbd8
const XK_leftshoe = 0xbda
const XK_lefttack = 0xbdc
const XK_righttack = 0xbfc

/* XK_APL */

/*
 * Hebrew
 * Byte 3 = c
 */

const XK_hebrew_doublelowline = 0xcdf
const XK_hebrew_aleph = 0xce0
const XK_hebrew_bet = 0xce1
const XK_hebrew_beth = 0xce1 /* deprecated */
const XK_hebrew_gimel = 0xce2
const XK_hebrew_gimmel = 0xce2 /* deprecated */
const XK_hebrew_dalet = 0xce3
const XK_hebrew_daleth = 0xce3 /* deprecated */
const XK_hebrew_he = 0xce4
const XK_hebrew_waw = 0xce5
const XK_hebrew_zain = 0xce6
const XK_hebrew_zayin = 0xce6 /* deprecated */
const XK_hebrew_chet = 0xce7
const XK_hebrew_het = 0xce7 /* deprecated */
const XK_hebrew_tet = 0xce8
const XK_hebrew_teth = 0xce8 /* deprecated */
const XK_hebrew_yod = 0xce9
const XK_hebrew_finalkaph = 0xcea
const XK_hebrew_kaph = 0xceb
const XK_hebrew_lamed = 0xcec
const XK_hebrew_finalmem = 0xced
const XK_hebrew_mem = 0xcee
const XK_hebrew_finalnun = 0xcef
const XK_hebrew_nun = 0xcf0
const XK_hebrew_samech = 0xcf1
const XK_hebrew_samekh = 0xcf1 /* deprecated */
const XK_hebrew_ayin = 0xcf2
const XK_hebrew_finalpe = 0xcf3
const XK_hebrew_pe = 0xcf4
const XK_hebrew_finalzade = 0xcf5
const XK_hebrew_finalzadi = 0xcf5 /* deprecated */
const XK_hebrew_zade = 0xcf6
const XK_hebrew_zadi = 0xcf6 /* deprecated */
const XK_hebrew_qoph = 0xcf7
const XK_hebrew_kuf = 0xcf7 /* deprecated */
const XK_hebrew_resh = 0xcf8
const XK_hebrew_shin = 0xcf9
const XK_hebrew_taw = 0xcfa
const XK_hebrew_taf = 0xcfa     /* deprecated */
const XK_Hebrew_switch = 0xFF7E /* Alias for mode_switch */
/* XK_HEBREW */

/*
 * Thai
 * Byte 3 = d
 */
const XK_Thai_kokai = 0xda1
const XK_Thai_khokhai = 0xda2
const XK_Thai_khokhuat = 0xda3
const XK_Thai_khokhwai = 0xda4
const XK_Thai_khokhon = 0xda5
const XK_Thai_khorakhang = 0xda6
const XK_Thai_ngongu = 0xda7
const XK_Thai_chochan = 0xda8
const XK_Thai_choching = 0xda9
const XK_Thai_chochang = 0xdaa
const XK_Thai_soso = 0xdab
const XK_Thai_chochoe = 0xdac
const XK_Thai_yoying = 0xdad
const XK_Thai_dochada = 0xdae
const XK_Thai_topatak = 0xdaf
const XK_Thai_thothan = 0xdb0
const XK_Thai_thonangmontho = 0xdb1
const XK_Thai_thophuthao = 0xdb2
const XK_Thai_nonen = 0xdb3
const XK_Thai_dodek = 0xdb4
const XK_Thai_totao = 0xdb5
const XK_Thai_thothung = 0xdb6
const XK_Thai_thothahan = 0xdb7
const XK_Thai_thothong = 0xdb8
const XK_Thai_nonu = 0xdb9
const XK_Thai_bobaimai = 0xdba
const XK_Thai_popla = 0xdbb
const XK_Thai_phophung = 0xdbc
const XK_Thai_fofa = 0xdbd
const XK_Thai_phophan = 0xdbe
const XK_Thai_fofan = 0xdbf
const XK_Thai_phosamphao = 0xdc0
const XK_Thai_moma = 0xdc1
const XK_Thai_yoyak = 0xdc2
const XK_Thai_rorua = 0xdc3
const XK_Thai_ru = 0xdc4
const XK_Thai_loling = 0xdc5
const XK_Thai_lu = 0xdc6
const XK_Thai_wowaen = 0xdc7
const XK_Thai_sosala = 0xdc8
const XK_Thai_sorusi = 0xdc9
const XK_Thai_sosua = 0xdca
const XK_Thai_hohip = 0xdcb
const XK_Thai_lochula = 0xdcc
const XK_Thai_oang = 0xdcd
const XK_Thai_honokhuk = 0xdce
const XK_Thai_paiyannoi = 0xdcf
const XK_Thai_saraa = 0xdd0
const XK_Thai_maihanakat = 0xdd1
const XK_Thai_saraaa = 0xdd2
const XK_Thai_saraam = 0xdd3
const XK_Thai_sarai = 0xdd4
const XK_Thai_saraii = 0xdd5
const XK_Thai_saraue = 0xdd6
const XK_Thai_sarauee = 0xdd7
const XK_Thai_sarau = 0xdd8
const XK_Thai_sarauu = 0xdd9
const XK_Thai_phinthu = 0xdda
const XK_Thai_maihanakat_maitho = 0xdde
const XK_Thai_baht = 0xddf
const XK_Thai_sarae = 0xde0
const XK_Thai_saraae = 0xde1
const XK_Thai_sarao = 0xde2
const XK_Thai_saraaimaimuan = 0xde3
const XK_Thai_saraaimaimalai = 0xde4
const XK_Thai_lakkhangyao = 0xde5
const XK_Thai_maiyamok = 0xde6
const XK_Thai_maitaikhu = 0xde7
const XK_Thai_maiek = 0xde8
const XK_Thai_maitho = 0xde9
const XK_Thai_maitri = 0xdea
const XK_Thai_maichattawa = 0xdeb
const XK_Thai_thanthakhat = 0xdec
const XK_Thai_nikhahit = 0xded
const XK_Thai_leksun = 0xdf0
const XK_Thai_leknung = 0xdf1
const XK_Thai_leksong = 0xdf2
const XK_Thai_leksam = 0xdf3
const XK_Thai_leksi = 0xdf4
const XK_Thai_lekha = 0xdf5
const XK_Thai_lekhok = 0xdf6
const XK_Thai_lekchet = 0xdf7
const XK_Thai_lekpaet = 0xdf8
const XK_Thai_lekkao = 0xdf9

/* XK_THAI */

/*
 *   Korean
 *   Byte 3 = e
 */
const XK_Hangul = 0xff31                   /* Hangul start/stop(toggle) */
const XK_Hangul_Start = 0xff32             /* Hangul start */
const XK_Hangul_End = 0xff33               /* Hangul end, English start */
const XK_Hangul_Hanja = 0xff34             /* Start Hangul->Hanja Conversion */
const XK_Hangul_Jamo = 0xff35              /* Hangul Jamo mode */
const XK_Hangul_Romaja = 0xff36            /* Hangul Romaja mode */
const XK_Hangul_Codeinput = 0xff37         /* Hangul code input mode */
const XK_Hangul_Jeonja = 0xff38            /* Jeonja mode */
const XK_Hangul_Banja = 0xff39             /* Banja mode */
const XK_Hangul_PreHanja = 0xff3a          /* Pre Hanja conversion */
const XK_Hangul_PostHanja = 0xff3b         /* Post Hanja conversion */
const XK_Hangul_SingleCandidate = 0xff3c   /* Single candidate */
const XK_Hangul_MultipleCandidate = 0xff3d /* Multiple candidate */
const XK_Hangul_PreviousCandidate = 0xff3e /* Previous candidate */
const XK_Hangul_Special = 0xff3f           /* Special symbols */
const XK_Hangul_switch = 0xFF7E            /* Alias for mode_switch */

/* Hangul Consonant Characters */
const XK_Hangul_Kiyeog = 0xea1
const XK_Hangul_SsangKiyeog = 0xea2
const XK_Hangul_KiyeogSios = 0xea3
const XK_Hangul_Nieun = 0xea4
const XK_Hangul_NieunJieuj = 0xea5
const XK_Hangul_NieunHieuh = 0xea6
const XK_Hangul_Dikeud = 0xea7
const XK_Hangul_SsangDikeud = 0xea8
const XK_Hangul_Rieul = 0xea9
const XK_Hangul_RieulKiyeog = 0xeaa
const XK_Hangul_RieulMieum = 0xeab
const XK_Hangul_RieulPieub = 0xeac
const XK_Hangul_RieulSios = 0xead
const XK_Hangul_RieulTieut = 0xeae
const XK_Hangul_RieulPhieuf = 0xeaf
const XK_Hangul_RieulHieuh = 0xeb0
const XK_Hangul_Mieum = 0xeb1
const XK_Hangul_Pieub = 0xeb2
const XK_Hangul_SsangPieub = 0xeb3
const XK_Hangul_PieubSios = 0xeb4
const XK_Hangul_Sios = 0xeb5
const XK_Hangul_SsangSios = 0xeb6
const XK_Hangul_Ieung = 0xeb7
const XK_Hangul_Jieuj = 0xeb8
const XK_Hangul_SsangJieuj = 0xeb9
const XK_Hangul_Cieuc = 0xeba
const XK_Hangul_Khieuq = 0xebb
const XK_Hangul_Tieut = 0xebc
const XK_Hangul_Phieuf = 0xebd
const XK_Hangul_Hieuh = 0xebe

/* Hangul Vowel Characters */
const XK_Hangul_A = 0xebf
const XK_Hangul_AE = 0xec0
const XK_Hangul_YA = 0xec1
const XK_Hangul_YAE = 0xec2
const XK_Hangul_EO = 0xec3
const XK_Hangul_E = 0xec4
const XK_Hangul_YEO = 0xec5
const XK_Hangul_YE = 0xec6
const XK_Hangul_O = 0xec7
const XK_Hangul_WA = 0xec8
const XK_Hangul_WAE = 0xec9
const XK_Hangul_OE = 0xeca
const XK_Hangul_YO = 0xecb
const XK_Hangul_U = 0xecc
const XK_Hangul_WEO = 0xecd
const XK_Hangul_WE = 0xece
const XK_Hangul_WI = 0xecf
const XK_Hangul_YU = 0xed0
const XK_Hangul_EU = 0xed1
const XK_Hangul_YI = 0xed2
const XK_Hangul_I = 0xed3

/* Hangul syllable-final (JongSeong) Characters */
const XK_Hangul_J_Kiyeog = 0xed4
const XK_Hangul_J_SsangKiyeog = 0xed5
const XK_Hangul_J_KiyeogSios = 0xed6
const XK_Hangul_J_Nieun = 0xed7
const XK_Hangul_J_NieunJieuj = 0xed8
const XK_Hangul_J_NieunHieuh = 0xed9
const XK_Hangul_J_Dikeud = 0xeda
const XK_Hangul_J_Rieul = 0xedb
const XK_Hangul_J_RieulKiyeog = 0xedc
const XK_Hangul_J_RieulMieum = 0xedd
const XK_Hangul_J_RieulPieub = 0xede
const XK_Hangul_J_RieulSios = 0xedf
const XK_Hangul_J_RieulTieut = 0xee0
const XK_Hangul_J_RieulPhieuf = 0xee1
const XK_Hangul_J_RieulHieuh = 0xee2
const XK_Hangul_J_Mieum = 0xee3
const XK_Hangul_J_Pieub = 0xee4
const XK_Hangul_J_PieubSios = 0xee5
const XK_Hangul_J_Sios = 0xee6
const XK_Hangul_J_SsangSios = 0xee7
const XK_Hangul_J_Ieung = 0xee8
const XK_Hangul_J_Jieuj = 0xee9
const XK_Hangul_J_Cieuc = 0xeea
const XK_Hangul_J_Khieuq = 0xeeb
const XK_Hangul_J_Tieut = 0xeec
const XK_Hangul_J_Phieuf = 0xeed
const XK_Hangul_J_Hieuh = 0xeee

/* Ancient Hangul Consonant Characters */
const XK_Hangul_RieulYeorinHieuh = 0xeef
const XK_Hangul_SunkyeongeumMieum = 0xef0
const XK_Hangul_SunkyeongeumPieub = 0xef1
const XK_Hangul_PanSios = 0xef2
const XK_Hangul_KkogjiDalrinIeung = 0xef3
const XK_Hangul_SunkyeongeumPhieuf = 0xef4
const XK_Hangul_YeorinHieuh = 0xef5

/* Ancient Hangul Vowel Characters */
const XK_Hangul_AraeA = 0xef6
const XK_Hangul_AraeAE = 0xef7

/* Ancient Hangul syllable-final (JongSeong) Characters */
const XK_Hangul_J_PanSios = 0xef8
const XK_Hangul_J_KkogjiDalrinIeung = 0xef9
const XK_Hangul_J_YeorinHieuh = 0xefa

/* Korean currency symbol */
const XK_Korean_Won = 0xeff

/* XK_KOREAN */

/* Euro currency symbol */
const XK_EuroSign = 0x20ac
