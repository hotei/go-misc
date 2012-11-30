// +build ignore

package go6

// #include <u.h>
// #include <6.out.h>
import "C"

const (
	AXXX             = C.AXXX
	AAAA             = C.AAAA
	AAAD             = C.AAAD
	AAAM             = C.AAAM
	AAAS             = C.AAAS
	AADCB            = C.AADCB
	AADCL            = C.AADCL
	AADCW            = C.AADCW
	AADDB            = C.AADDB
	AADDL            = C.AADDL
	AADDW            = C.AADDW
	AADJSP           = C.AADJSP
	AANDB            = C.AANDB
	AANDL            = C.AANDL
	AANDW            = C.AANDW
	AARPL            = C.AARPL
	ABOUNDL          = C.ABOUNDL
	ABOUNDW          = C.ABOUNDW
	ABSFL            = C.ABSFL
	ABSFW            = C.ABSFW
	ABSRL            = C.ABSRL
	ABSRW            = C.ABSRW
	ABTL             = C.ABTL
	ABTW             = C.ABTW
	ABTCL            = C.ABTCL
	ABTCW            = C.ABTCW
	ABTRL            = C.ABTRL
	ABTRW            = C.ABTRW
	ABTSL            = C.ABTSL
	ABTSW            = C.ABTSW
	ABYTE            = C.ABYTE
	ACALL            = C.ACALL
	ACLC             = C.ACLC
	ACLD             = C.ACLD
	ACLI             = C.ACLI
	ACLTS            = C.ACLTS
	ACMC             = C.ACMC
	ACMPB            = C.ACMPB
	ACMPL            = C.ACMPL
	ACMPW            = C.ACMPW
	ACMPSB           = C.ACMPSB
	ACMPSL           = C.ACMPSL
	ACMPSW           = C.ACMPSW
	ADAA             = C.ADAA
	ADAS             = C.ADAS
	ADATA            = C.ADATA
	ADECB            = C.ADECB
	ADECL            = C.ADECL
	ADECQ            = C.ADECQ
	ADECW            = C.ADECW
	ADIVB            = C.ADIVB
	ADIVL            = C.ADIVL
	ADIVW            = C.ADIVW
	AENTER           = C.AENTER
	AGLOBL           = C.AGLOBL
	AGOK             = C.AGOK
	AHISTORY         = C.AHISTORY
	AHLT             = C.AHLT
	AIDIVB           = C.AIDIVB
	AIDIVL           = C.AIDIVL
	AIDIVW           = C.AIDIVW
	AIMULB           = C.AIMULB
	AIMULL           = C.AIMULL
	AIMULW           = C.AIMULW
	AINB             = C.AINB
	AINL             = C.AINL
	AINW             = C.AINW
	AINCB            = C.AINCB
	AINCL            = C.AINCL
	AINCQ            = C.AINCQ
	AINCW            = C.AINCW
	AINSB            = C.AINSB
	AINSL            = C.AINSL
	AINSW            = C.AINSW
	AINT             = C.AINT
	AINTO            = C.AINTO
	AIRETL           = C.AIRETL
	AIRETW           = C.AIRETW
	AJCC             = C.AJCC
	AJCS             = C.AJCS
	AJCXZL           = C.AJCXZL
	AJEQ             = C.AJEQ
	AJGE             = C.AJGE
	AJGT             = C.AJGT
	AJHI             = C.AJHI
	AJLE             = C.AJLE
	AJLS             = C.AJLS
	AJLT             = C.AJLT
	AJMI             = C.AJMI
	AJMP             = C.AJMP
	AJNE             = C.AJNE
	AJOC             = C.AJOC
	AJOS             = C.AJOS
	AJPC             = C.AJPC
	AJPL             = C.AJPL
	AJPS             = C.AJPS
	ALAHF            = C.ALAHF
	ALARL            = C.ALARL
	ALARW            = C.ALARW
	ALEAL            = C.ALEAL
	ALEAW            = C.ALEAW
	ALEAVEL          = C.ALEAVEL
	ALEAVEW          = C.ALEAVEW
	ALOCK            = C.ALOCK
	ALODSB           = C.ALODSB
	ALODSL           = C.ALODSL
	ALODSW           = C.ALODSW
	ALONG            = C.ALONG
	ALOOP            = C.ALOOP
	ALOOPEQ          = C.ALOOPEQ
	ALOOPNE          = C.ALOOPNE
	ALSLL            = C.ALSLL
	ALSLW            = C.ALSLW
	AMOVB            = C.AMOVB
	AMOVL            = C.AMOVL
	AMOVW            = C.AMOVW
	AMOVBLSX         = C.AMOVBLSX
	AMOVBLZX         = C.AMOVBLZX
	AMOVBQSX         = C.AMOVBQSX
	AMOVBQZX         = C.AMOVBQZX
	AMOVBWSX         = C.AMOVBWSX
	AMOVBWZX         = C.AMOVBWZX
	AMOVWLSX         = C.AMOVWLSX
	AMOVWLZX         = C.AMOVWLZX
	AMOVWQSX         = C.AMOVWQSX
	AMOVWQZX         = C.AMOVWQZX
	AMOVSB           = C.AMOVSB
	AMOVSL           = C.AMOVSL
	AMOVSW           = C.AMOVSW
	AMULB            = C.AMULB
	AMULL            = C.AMULL
	AMULW            = C.AMULW
	ANAME            = C.ANAME
	ANEGB            = C.ANEGB
	ANEGL            = C.ANEGL
	ANEGW            = C.ANEGW
	ANOP             = C.ANOP
	ANOTB            = C.ANOTB
	ANOTL            = C.ANOTL
	ANOTW            = C.ANOTW
	AORB             = C.AORB
	AORL             = C.AORL
	AORW             = C.AORW
	AOUTB            = C.AOUTB
	AOUTL            = C.AOUTL
	AOUTW            = C.AOUTW
	AOUTSB           = C.AOUTSB
	AOUTSL           = C.AOUTSL
	AOUTSW           = C.AOUTSW
	APAUSE           = C.APAUSE
	APOPAL           = C.APOPAL
	APOPAW           = C.APOPAW
	APOPFL           = C.APOPFL
	APOPFW           = C.APOPFW
	APOPL            = C.APOPL
	APOPW            = C.APOPW
	APUSHAL          = C.APUSHAL
	APUSHAW          = C.APUSHAW
	APUSHFL          = C.APUSHFL
	APUSHFW          = C.APUSHFW
	APUSHL           = C.APUSHL
	APUSHW           = C.APUSHW
	ARCLB            = C.ARCLB
	ARCLL            = C.ARCLL
	ARCLW            = C.ARCLW
	ARCRB            = C.ARCRB
	ARCRL            = C.ARCRL
	ARCRW            = C.ARCRW
	AREP             = C.AREP
	AREPN            = C.AREPN
	ARET             = C.ARET
	AROLB            = C.AROLB
	AROLL            = C.AROLL
	AROLW            = C.AROLW
	ARORB            = C.ARORB
	ARORL            = C.ARORL
	ARORW            = C.ARORW
	ASAHF            = C.ASAHF
	ASALB            = C.ASALB
	ASALL            = C.ASALL
	ASALW            = C.ASALW
	ASARB            = C.ASARB
	ASARL            = C.ASARL
	ASARW            = C.ASARW
	ASBBB            = C.ASBBB
	ASBBL            = C.ASBBL
	ASBBW            = C.ASBBW
	ASCASB           = C.ASCASB
	ASCASL           = C.ASCASL
	ASCASW           = C.ASCASW
	ASETCC           = C.ASETCC
	ASETCS           = C.ASETCS
	ASETEQ           = C.ASETEQ
	ASETGE           = C.ASETGE
	ASETGT           = C.ASETGT
	ASETHI           = C.ASETHI
	ASETLE           = C.ASETLE
	ASETLS           = C.ASETLS
	ASETLT           = C.ASETLT
	ASETMI           = C.ASETMI
	ASETNE           = C.ASETNE
	ASETOC           = C.ASETOC
	ASETOS           = C.ASETOS
	ASETPC           = C.ASETPC
	ASETPL           = C.ASETPL
	ASETPS           = C.ASETPS
	ACDQ             = C.ACDQ
	ACWD             = C.ACWD
	ASHLB            = C.ASHLB
	ASHLL            = C.ASHLL
	ASHLW            = C.ASHLW
	ASHRB            = C.ASHRB
	ASHRL            = C.ASHRL
	ASHRW            = C.ASHRW
	ASTC             = C.ASTC
	ASTD             = C.ASTD
	ASTI             = C.ASTI
	ASTOSB           = C.ASTOSB
	ASTOSL           = C.ASTOSL
	ASTOSW           = C.ASTOSW
	ASUBB            = C.ASUBB
	ASUBL            = C.ASUBL
	ASUBW            = C.ASUBW
	ASYSCALL         = C.ASYSCALL
	ATESTB           = C.ATESTB
	ATESTL           = C.ATESTL
	ATESTW           = C.ATESTW
	ATEXT            = C.ATEXT
	AVERR            = C.AVERR
	AVERW            = C.AVERW
	AWAIT            = C.AWAIT
	AWORD            = C.AWORD
	AXCHGB           = C.AXCHGB
	AXCHGL           = C.AXCHGL
	AXCHGW           = C.AXCHGW
	AXLAT            = C.AXLAT
	AXORB            = C.AXORB
	AXORL            = C.AXORL
	AXORW            = C.AXORW
	AFMOVB           = C.AFMOVB
	AFMOVBP          = C.AFMOVBP
	AFMOVD           = C.AFMOVD
	AFMOVDP          = C.AFMOVDP
	AFMOVF           = C.AFMOVF
	AFMOVFP          = C.AFMOVFP
	AFMOVL           = C.AFMOVL
	AFMOVLP          = C.AFMOVLP
	AFMOVV           = C.AFMOVV
	AFMOVVP          = C.AFMOVVP
	AFMOVW           = C.AFMOVW
	AFMOVWP          = C.AFMOVWP
	AFMOVX           = C.AFMOVX
	AFMOVXP          = C.AFMOVXP
	AFCOMB           = C.AFCOMB
	AFCOMBP          = C.AFCOMBP
	AFCOMD           = C.AFCOMD
	AFCOMDP          = C.AFCOMDP
	AFCOMDPP         = C.AFCOMDPP
	AFCOMF           = C.AFCOMF
	AFCOMFP          = C.AFCOMFP
	AFCOML           = C.AFCOML
	AFCOMLP          = C.AFCOMLP
	AFCOMW           = C.AFCOMW
	AFCOMWP          = C.AFCOMWP
	AFUCOM           = C.AFUCOM
	AFUCOMP          = C.AFUCOMP
	AFUCOMPP         = C.AFUCOMPP
	AFADDDP          = C.AFADDDP
	AFADDW           = C.AFADDW
	AFADDL           = C.AFADDL
	AFADDF           = C.AFADDF
	AFADDD           = C.AFADDD
	AFMULDP          = C.AFMULDP
	AFMULW           = C.AFMULW
	AFMULL           = C.AFMULL
	AFMULF           = C.AFMULF
	AFMULD           = C.AFMULD
	AFSUBDP          = C.AFSUBDP
	AFSUBW           = C.AFSUBW
	AFSUBL           = C.AFSUBL
	AFSUBF           = C.AFSUBF
	AFSUBD           = C.AFSUBD
	AFSUBRDP         = C.AFSUBRDP
	AFSUBRW          = C.AFSUBRW
	AFSUBRL          = C.AFSUBRL
	AFSUBRF          = C.AFSUBRF
	AFSUBRD          = C.AFSUBRD
	AFDIVDP          = C.AFDIVDP
	AFDIVW           = C.AFDIVW
	AFDIVL           = C.AFDIVL
	AFDIVF           = C.AFDIVF
	AFDIVD           = C.AFDIVD
	AFDIVRDP         = C.AFDIVRDP
	AFDIVRW          = C.AFDIVRW
	AFDIVRL          = C.AFDIVRL
	AFDIVRF          = C.AFDIVRF
	AFDIVRD          = C.AFDIVRD
	AFXCHD           = C.AFXCHD
	AFFREE           = C.AFFREE
	AFLDCW           = C.AFLDCW
	AFLDENV          = C.AFLDENV
	AFRSTOR          = C.AFRSTOR
	AFSAVE           = C.AFSAVE
	AFSTCW           = C.AFSTCW
	AFSTENV          = C.AFSTENV
	AFSTSW           = C.AFSTSW
	AF2XM1           = C.AF2XM1
	AFABS            = C.AFABS
	AFCHS            = C.AFCHS
	AFCLEX           = C.AFCLEX
	AFCOS            = C.AFCOS
	AFDECSTP         = C.AFDECSTP
	AFINCSTP         = C.AFINCSTP
	AFINIT           = C.AFINIT
	AFLD1            = C.AFLD1
	AFLDL2E          = C.AFLDL2E
	AFLDL2T          = C.AFLDL2T
	AFLDLG2          = C.AFLDLG2
	AFLDLN2          = C.AFLDLN2
	AFLDPI           = C.AFLDPI
	AFLDZ            = C.AFLDZ
	AFNOP            = C.AFNOP
	AFPATAN          = C.AFPATAN
	AFPREM           = C.AFPREM
	AFPREM1          = C.AFPREM1
	AFPTAN           = C.AFPTAN
	AFRNDINT         = C.AFRNDINT
	AFSCALE          = C.AFSCALE
	AFSIN            = C.AFSIN
	AFSINCOS         = C.AFSINCOS
	AFSQRT           = C.AFSQRT
	AFTST            = C.AFTST
	AFXAM            = C.AFXAM
	AFXTRACT         = C.AFXTRACT
	AFYL2X           = C.AFYL2X
	AFYL2XP1         = C.AFYL2XP1
	AEND             = C.AEND
	ADYNT_           = C.ADYNT_
	AINIT_           = C.AINIT_
	ASIGNAME         = C.ASIGNAME
	ACMPXCHGB        = C.ACMPXCHGB
	ACMPXCHGL        = C.ACMPXCHGL
	ACMPXCHGW        = C.ACMPXCHGW
	ACMPXCHG8B       = C.ACMPXCHG8B
	ACPUID           = C.ACPUID
	AINVD            = C.AINVD
	AINVLPG          = C.AINVLPG
	ALFENCE          = C.ALFENCE
	AMFENCE          = C.AMFENCE
	AMOVNTIL         = C.AMOVNTIL
	ARDMSR           = C.ARDMSR
	ARDPMC           = C.ARDPMC
	ARDTSC           = C.ARDTSC
	ARSM             = C.ARSM
	ASFENCE          = C.ASFENCE
	ASYSRET          = C.ASYSRET
	AWBINVD          = C.AWBINVD
	AWRMSR           = C.AWRMSR
	AXADDB           = C.AXADDB
	AXADDL           = C.AXADDL
	AXADDW           = C.AXADDW
	ACMOVLCC         = C.ACMOVLCC
	ACMOVLCS         = C.ACMOVLCS
	ACMOVLEQ         = C.ACMOVLEQ
	ACMOVLGE         = C.ACMOVLGE
	ACMOVLGT         = C.ACMOVLGT
	ACMOVLHI         = C.ACMOVLHI
	ACMOVLLE         = C.ACMOVLLE
	ACMOVLLS         = C.ACMOVLLS
	ACMOVLLT         = C.ACMOVLLT
	ACMOVLMI         = C.ACMOVLMI
	ACMOVLNE         = C.ACMOVLNE
	ACMOVLOC         = C.ACMOVLOC
	ACMOVLOS         = C.ACMOVLOS
	ACMOVLPC         = C.ACMOVLPC
	ACMOVLPL         = C.ACMOVLPL
	ACMOVLPS         = C.ACMOVLPS
	ACMOVQCC         = C.ACMOVQCC
	ACMOVQCS         = C.ACMOVQCS
	ACMOVQEQ         = C.ACMOVQEQ
	ACMOVQGE         = C.ACMOVQGE
	ACMOVQGT         = C.ACMOVQGT
	ACMOVQHI         = C.ACMOVQHI
	ACMOVQLE         = C.ACMOVQLE
	ACMOVQLS         = C.ACMOVQLS
	ACMOVQLT         = C.ACMOVQLT
	ACMOVQMI         = C.ACMOVQMI
	ACMOVQNE         = C.ACMOVQNE
	ACMOVQOC         = C.ACMOVQOC
	ACMOVQOS         = C.ACMOVQOS
	ACMOVQPC         = C.ACMOVQPC
	ACMOVQPL         = C.ACMOVQPL
	ACMOVQPS         = C.ACMOVQPS
	ACMOVWCC         = C.ACMOVWCC
	ACMOVWCS         = C.ACMOVWCS
	ACMOVWEQ         = C.ACMOVWEQ
	ACMOVWGE         = C.ACMOVWGE
	ACMOVWGT         = C.ACMOVWGT
	ACMOVWHI         = C.ACMOVWHI
	ACMOVWLE         = C.ACMOVWLE
	ACMOVWLS         = C.ACMOVWLS
	ACMOVWLT         = C.ACMOVWLT
	ACMOVWMI         = C.ACMOVWMI
	ACMOVWNE         = C.ACMOVWNE
	ACMOVWOC         = C.ACMOVWOC
	ACMOVWOS         = C.ACMOVWOS
	ACMOVWPC         = C.ACMOVWPC
	ACMOVWPL         = C.ACMOVWPL
	ACMOVWPS         = C.ACMOVWPS
	AADCQ            = C.AADCQ
	AADDQ            = C.AADDQ
	AANDQ            = C.AANDQ
	ABSFQ            = C.ABSFQ
	ABSRQ            = C.ABSRQ
	ABTCQ            = C.ABTCQ
	ABTQ             = C.ABTQ
	ABTRQ            = C.ABTRQ
	ABTSQ            = C.ABTSQ
	ACMPQ            = C.ACMPQ
	ACMPSQ           = C.ACMPSQ
	ACMPXCHGQ        = C.ACMPXCHGQ
	ACQO             = C.ACQO
	ADIVQ            = C.ADIVQ
	AIDIVQ           = C.AIDIVQ
	AIMULQ           = C.AIMULQ
	AIRETQ           = C.AIRETQ
	AJCXZQ           = C.AJCXZQ
	ALEAQ            = C.ALEAQ
	ALEAVEQ          = C.ALEAVEQ
	ALODSQ           = C.ALODSQ
	AMOVQ            = C.AMOVQ
	AMOVLQSX         = C.AMOVLQSX
	AMOVLQZX         = C.AMOVLQZX
	AMOVNTIQ         = C.AMOVNTIQ
	AMOVSQ           = C.AMOVSQ
	AMULQ            = C.AMULQ
	ANEGQ            = C.ANEGQ
	ANOTQ            = C.ANOTQ
	AORQ             = C.AORQ
	APOPFQ           = C.APOPFQ
	APOPQ            = C.APOPQ
	APUSHFQ          = C.APUSHFQ
	APUSHQ           = C.APUSHQ
	ARCLQ            = C.ARCLQ
	ARCRQ            = C.ARCRQ
	AROLQ            = C.AROLQ
	ARORQ            = C.ARORQ
	AQUAD            = C.AQUAD
	ASALQ            = C.ASALQ
	ASARQ            = C.ASARQ
	ASBBQ            = C.ASBBQ
	ASCASQ           = C.ASCASQ
	ASHLQ            = C.ASHLQ
	ASHRQ            = C.ASHRQ
	ASTOSQ           = C.ASTOSQ
	ASUBQ            = C.ASUBQ
	ATESTQ           = C.ATESTQ
	AXADDQ           = C.AXADDQ
	AXCHGQ           = C.AXCHGQ
	AXORQ            = C.AXORQ
	AADDPD           = C.AADDPD
	AADDPS           = C.AADDPS
	AADDSD           = C.AADDSD
	AADDSS           = C.AADDSS
	AANDNPD          = C.AANDNPD
	AANDNPS          = C.AANDNPS
	AANDPD           = C.AANDPD
	AANDPS           = C.AANDPS
	ACMPPD           = C.ACMPPD
	ACMPPS           = C.ACMPPS
	ACMPSD           = C.ACMPSD
	ACMPSS           = C.ACMPSS
	ACOMISD          = C.ACOMISD
	ACOMISS          = C.ACOMISS
	ACVTPD2PL        = C.ACVTPD2PL
	ACVTPD2PS        = C.ACVTPD2PS
	ACVTPL2PD        = C.ACVTPL2PD
	ACVTPL2PS        = C.ACVTPL2PS
	ACVTPS2PD        = C.ACVTPS2PD
	ACVTPS2PL        = C.ACVTPS2PL
	ACVTSD2SL        = C.ACVTSD2SL
	ACVTSD2SQ        = C.ACVTSD2SQ
	ACVTSD2SS        = C.ACVTSD2SS
	ACVTSL2SD        = C.ACVTSL2SD
	ACVTSL2SS        = C.ACVTSL2SS
	ACVTSQ2SD        = C.ACVTSQ2SD
	ACVTSQ2SS        = C.ACVTSQ2SS
	ACVTSS2SD        = C.ACVTSS2SD
	ACVTSS2SL        = C.ACVTSS2SL
	ACVTSS2SQ        = C.ACVTSS2SQ
	ACVTTPD2PL       = C.ACVTTPD2PL
	ACVTTPS2PL       = C.ACVTTPS2PL
	ACVTTSD2SL       = C.ACVTTSD2SL
	ACVTTSD2SQ       = C.ACVTTSD2SQ
	ACVTTSS2SL       = C.ACVTTSS2SL
	ACVTTSS2SQ       = C.ACVTTSS2SQ
	ADIVPD           = C.ADIVPD
	ADIVPS           = C.ADIVPS
	ADIVSD           = C.ADIVSD
	ADIVSS           = C.ADIVSS
	AEMMS            = C.AEMMS
	AFXRSTOR         = C.AFXRSTOR
	AFXRSTOR64       = C.AFXRSTOR64
	AFXSAVE          = C.AFXSAVE
	AFXSAVE64        = C.AFXSAVE64
	ALDMXCSR         = C.ALDMXCSR
	AMASKMOVOU       = C.AMASKMOVOU
	AMASKMOVQ        = C.AMASKMOVQ
	AMAXPD           = C.AMAXPD
	AMAXPS           = C.AMAXPS
	AMAXSD           = C.AMAXSD
	AMAXSS           = C.AMAXSS
	AMINPD           = C.AMINPD
	AMINPS           = C.AMINPS
	AMINSD           = C.AMINSD
	AMINSS           = C.AMINSS
	AMOVAPD          = C.AMOVAPD
	AMOVAPS          = C.AMOVAPS
	AMOVOU           = C.AMOVOU
	AMOVHLPS         = C.AMOVHLPS
	AMOVHPD          = C.AMOVHPD
	AMOVHPS          = C.AMOVHPS
	AMOVLHPS         = C.AMOVLHPS
	AMOVLPD          = C.AMOVLPD
	AMOVLPS          = C.AMOVLPS
	AMOVMSKPD        = C.AMOVMSKPD
	AMOVMSKPS        = C.AMOVMSKPS
	AMOVNTO          = C.AMOVNTO
	AMOVNTPD         = C.AMOVNTPD
	AMOVNTPS         = C.AMOVNTPS
	AMOVNTQ          = C.AMOVNTQ
	AMOVO            = C.AMOVO
	AMOVQOZX         = C.AMOVQOZX
	AMOVSD           = C.AMOVSD
	AMOVSS           = C.AMOVSS
	AMOVUPD          = C.AMOVUPD
	AMOVUPS          = C.AMOVUPS
	AMULPD           = C.AMULPD
	AMULPS           = C.AMULPS
	AMULSD           = C.AMULSD
	AMULSS           = C.AMULSS
	AORPD            = C.AORPD
	AORPS            = C.AORPS
	APACKSSLW        = C.APACKSSLW
	APACKSSWB        = C.APACKSSWB
	APACKUSWB        = C.APACKUSWB
	APADDB           = C.APADDB
	APADDL           = C.APADDL
	APADDQ           = C.APADDQ
	APADDSB          = C.APADDSB
	APADDSW          = C.APADDSW
	APADDUSB         = C.APADDUSB
	APADDUSW         = C.APADDUSW
	APADDW           = C.APADDW
	APANDB           = C.APANDB
	APANDL           = C.APANDL
	APANDSB          = C.APANDSB
	APANDSW          = C.APANDSW
	APANDUSB         = C.APANDUSB
	APANDUSW         = C.APANDUSW
	APANDW           = C.APANDW
	APAND            = C.APAND
	APANDN           = C.APANDN
	APAVGB           = C.APAVGB
	APAVGW           = C.APAVGW
	APCMPEQB         = C.APCMPEQB
	APCMPEQL         = C.APCMPEQL
	APCMPEQW         = C.APCMPEQW
	APCMPGTB         = C.APCMPGTB
	APCMPGTL         = C.APCMPGTL
	APCMPGTW         = C.APCMPGTW
	APEXTRW          = C.APEXTRW
	APFACC           = C.APFACC
	APFADD           = C.APFADD
	APFCMPEQ         = C.APFCMPEQ
	APFCMPGE         = C.APFCMPGE
	APFCMPGT         = C.APFCMPGT
	APFMAX           = C.APFMAX
	APFMIN           = C.APFMIN
	APFMUL           = C.APFMUL
	APFNACC          = C.APFNACC
	APFPNACC         = C.APFPNACC
	APFRCP           = C.APFRCP
	APFRCPIT1        = C.APFRCPIT1
	APFRCPI2T        = C.APFRCPI2T
	APFRSQIT1        = C.APFRSQIT1
	APFRSQRT         = C.APFRSQRT
	APFSUB           = C.APFSUB
	APFSUBR          = C.APFSUBR
	APINSRW          = C.APINSRW
	APMADDWL         = C.APMADDWL
	APMAXSW          = C.APMAXSW
	APMAXUB          = C.APMAXUB
	APMINSW          = C.APMINSW
	APMINUB          = C.APMINUB
	APMOVMSKB        = C.APMOVMSKB
	APMULHRW         = C.APMULHRW
	APMULHUW         = C.APMULHUW
	APMULHW          = C.APMULHW
	APMULLW          = C.APMULLW
	APMULULQ         = C.APMULULQ
	APOR             = C.APOR
	APSADBW          = C.APSADBW
	APSHUFHW         = C.APSHUFHW
	APSHUFL          = C.APSHUFL
	APSHUFLW         = C.APSHUFLW
	APSHUFW          = C.APSHUFW
	APSLLO           = C.APSLLO
	APSLLL           = C.APSLLL
	APSLLQ           = C.APSLLQ
	APSLLW           = C.APSLLW
	APSRAL           = C.APSRAL
	APSRAW           = C.APSRAW
	APSRLO           = C.APSRLO
	APSRLL           = C.APSRLL
	APSRLQ           = C.APSRLQ
	APSRLW           = C.APSRLW
	APSUBB           = C.APSUBB
	APSUBL           = C.APSUBL
	APSUBQ           = C.APSUBQ
	APSUBSB          = C.APSUBSB
	APSUBSW          = C.APSUBSW
	APSUBUSB         = C.APSUBUSB
	APSUBUSW         = C.APSUBUSW
	APSUBW           = C.APSUBW
	APSWAPL          = C.APSWAPL
	APUNPCKHBW       = C.APUNPCKHBW
	APUNPCKHLQ       = C.APUNPCKHLQ
	APUNPCKHQDQ      = C.APUNPCKHQDQ
	APUNPCKHWL       = C.APUNPCKHWL
	APUNPCKLBW       = C.APUNPCKLBW
	APUNPCKLLQ       = C.APUNPCKLLQ
	APUNPCKLQDQ      = C.APUNPCKLQDQ
	APUNPCKLWL       = C.APUNPCKLWL
	APXOR            = C.APXOR
	ARCPPS           = C.ARCPPS
	ARCPSS           = C.ARCPSS
	ARSQRTPS         = C.ARSQRTPS
	ARSQRTSS         = C.ARSQRTSS
	ASHUFPD          = C.ASHUFPD
	ASHUFPS          = C.ASHUFPS
	ASQRTPD          = C.ASQRTPD
	ASQRTPS          = C.ASQRTPS
	ASQRTSD          = C.ASQRTSD
	ASQRTSS          = C.ASQRTSS
	ASTMXCSR         = C.ASTMXCSR
	ASUBPD           = C.ASUBPD
	ASUBPS           = C.ASUBPS
	ASUBSD           = C.ASUBSD
	ASUBSS           = C.ASUBSS
	AUCOMISD         = C.AUCOMISD
	AUCOMISS         = C.AUCOMISS
	AUNPCKHPD        = C.AUNPCKHPD
	AUNPCKHPS        = C.AUNPCKHPS
	AUNPCKLPD        = C.AUNPCKLPD
	AUNPCKLPS        = C.AUNPCKLPS
	AXORPD           = C.AXORPD
	AXORPS           = C.AXORPS
	APF2IW           = C.APF2IW
	APF2IL           = C.APF2IL
	API2FW           = C.API2FW
	API2FL           = C.API2FL
	ARETFW           = C.ARETFW
	ARETFL           = C.ARETFL
	ARETFQ           = C.ARETFQ
	ASWAPGS          = C.ASWAPGS
	AMODE            = C.AMODE
	ACRC32B          = C.ACRC32B
	ACRC32Q          = C.ACRC32Q
	AIMUL3Q          = C.AIMUL3Q
	APREFETCHT0      = C.APREFETCHT0
	APREFETCHT1      = C.APREFETCHT1
	APREFETCHT2      = C.APREFETCHT2
	APREFETCHNTA     = C.APREFETCHNTA
	AMOVQL           = C.AMOVQL
	ABSWAPL          = C.ABSWAPL
	ABSWAPQ          = C.ABSWAPQ
	AUNDEF           = C.AUNDEF
	AAESENC          = C.AAESENC
	AAESENCLAST      = C.AAESENCLAST
	AAESDEC          = C.AAESDEC
	AAESDECLAST      = C.AAESDECLAST
	AAESIMC          = C.AAESIMC
	AAESKEYGENASSIST = C.AAESKEYGENASSIST
	APSHUFD          = C.APSHUFD
	AUSEFIELD        = C.AUSEFIELD
	ALAST            = C.ALAST
	D_AL             = C.D_AL
	D_CL             = C.D_CL
	D_DL             = C.D_DL
	D_BL             = C.D_BL
	D_SPB            = C.D_SPB
	D_BPB            = C.D_BPB
	D_SIB            = C.D_SIB
	D_DIB            = C.D_DIB
	D_R8B            = C.D_R8B
	D_R9B            = C.D_R9B
	D_R10B           = C.D_R10B
	D_R11B           = C.D_R11B
	D_R12B           = C.D_R12B
	D_R13B           = C.D_R13B
	D_R14B           = C.D_R14B
	D_R15B           = C.D_R15B
	D_AX             = C.D_AX
	D_CX             = C.D_CX
	D_DX             = C.D_DX
	D_BX             = C.D_BX
	D_SP             = C.D_SP
	D_BP             = C.D_BP
	D_SI             = C.D_SI
	D_DI             = C.D_DI
	D_R8             = C.D_R8
	D_R9             = C.D_R9
	D_R10            = C.D_R10
	D_R11            = C.D_R11
	D_R12            = C.D_R12
	D_R13            = C.D_R13
	D_R14            = C.D_R14
	D_R15            = C.D_R15
	D_AH             = C.D_AH
	D_CH             = C.D_CH
	D_DH             = C.D_DH
	D_BH             = C.D_BH
	D_F0             = C.D_F0
	D_M0             = C.D_M0
	D_X0             = C.D_X0
	D_X1             = C.D_X1
	D_X2             = C.D_X2
	D_X3             = C.D_X3
	D_X4             = C.D_X4
	D_X5             = C.D_X5
	D_X6             = C.D_X6
	D_X7             = C.D_X7
	D_X8             = C.D_X8
	D_X9             = C.D_X9
	D_X10            = C.D_X10
	D_X11            = C.D_X11
	D_X12            = C.D_X12
	D_X13            = C.D_X13
	D_X14            = C.D_X14
	D_X15            = C.D_X15
	D_CS             = C.D_CS
	D_SS             = C.D_SS
	D_DS             = C.D_DS
	D_ES             = C.D_ES
	D_FS             = C.D_FS
	D_GS             = C.D_GS
	D_GDTR           = C.D_GDTR
	D_IDTR           = C.D_IDTR
	D_LDTR           = C.D_LDTR
	D_MSW            = C.D_MSW
	D_TASK           = C.D_TASK
	D_CR             = C.D_CR
	D_DR             = C.D_DR
	D_TR             = C.D_TR
	D_NONE           = C.D_NONE
	D_BRANCH         = C.D_BRANCH
	D_EXTERN         = C.D_EXTERN
	D_STATIC         = C.D_STATIC
	D_AUTO           = C.D_AUTO
	D_PARAM          = C.D_PARAM
	D_CONST          = C.D_CONST
	D_FCONST         = C.D_FCONST
	D_SCONST         = C.D_SCONST
	D_ADDR           = C.D_ADDR
	D_FILE           = C.D_FILE
	D_FILE1          = C.D_FILE1
	D_INDIR          = C.D_INDIR
	D_SIZE           = C.D_SIZE
	D_PCREL          = C.D_PCREL
	T_TYPE           = C.T_TYPE
	T_INDEX          = C.T_INDEX
	T_OFFSET         = C.T_OFFSET
	T_FCONST         = C.T_FCONST
	T_SYM            = C.T_SYM
	T_SCONST         = C.T_SCONST
	T_64             = C.T_64
	T_GOTYPE         = C.T_GOTYPE
)
