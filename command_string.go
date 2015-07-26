// generated by stringer -type Command command.go; DO NOT EDIT

package ipmsg

import "fmt"

const _Command_name = "NOOPERATIONBR_ENTRYBR_EXITANSENTRYBR_ABSENCEBR_ISGETLISTOKGETLISTGETLISTANSLISTBR_ISGETLIST2SENDMSGRECVMSGREADMSGDELMSGANSREADMSGGETINFOSENDINFOGETABSENCEINFOSENDABSENCEINFOGETFILEDATRELEASEFILGETDIRFILEGETPUBKEYANSPUBKEYMODESENDCHECKSECRETBROADCASTMULTICASTNOPOPUPAUTORETRETRYPASSWORDNOLOGNEWMUTINOADDLISTREADCHECKFILEATTACHENCRYPT"

var _Command_map = map[Command]string{
	0:       _Command_name[0:11],
	1:       _Command_name[11:19],
	2:       _Command_name[19:26],
	3:       _Command_name[26:34],
	4:       _Command_name[34:44],
	16:      _Command_name[44:56],
	17:      _Command_name[56:65],
	18:      _Command_name[65:72],
	19:      _Command_name[72:79],
	24:      _Command_name[79:92],
	32:      _Command_name[92:99],
	33:      _Command_name[99:106],
	48:      _Command_name[106:113],
	49:      _Command_name[113:119],
	50:      _Command_name[119:129],
	64:      _Command_name[129:136],
	65:      _Command_name[136:144],
	80:      _Command_name[144:158],
	81:      _Command_name[158:173],
	96:      _Command_name[173:183],
	97:      _Command_name[183:193],
	98:      _Command_name[193:203],
	114:     _Command_name[203:212],
	115:     _Command_name[212:221],
	255:     _Command_name[221:225],
	256:     _Command_name[225:234],
	512:     _Command_name[234:240],
	1024:    _Command_name[240:249],
	2048:    _Command_name[249:258],
	4096:    _Command_name[258:265],
	8192:    _Command_name[265:272],
	16384:   _Command_name[272:277],
	32768:   _Command_name[277:285],
	131072:  _Command_name[285:290],
	262144:  _Command_name[290:297],
	524288:  _Command_name[297:306],
	1048576: _Command_name[306:315],
	2097152: _Command_name[315:325],
	4194304: _Command_name[325:332],
}

func (i Command) String() string {
	if str, ok := _Command_map[i]; ok {
		return str
	}
	return fmt.Sprintf("Command(%d)", i)
}