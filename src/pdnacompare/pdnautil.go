package pdnacompare

import (
	"fmt"
	//"time"
	//"log"
	"runtime"
)

const PDNALEN = 144
const MAXDISTANCE = 15001
const SHORTDISTANCE = 150

// this is 10-15% faster than having a loop
func pdnaCompareStatic(ref_dna []byte, reporDNA []byte) int {
	//var maxTotal = 2000
	var runnintTotal uint16
	var x uint8

	//var r,p uint8


	//r,p = uint8(ref_dna[0]),uint8(reporDNA[0]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
	/*	r,p = uint8(ref_dna[1]),uint8(reporDNA[1]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[2]),uint8(reporDNA[2]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[3]),uint8(reporDNA[3]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[4]),uint8(reporDNA[4]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[5]),uint8(reporDNA[5]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[6]),uint8(reporDNA[6]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[7]),uint8(reporDNA[7]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[8]),uint8(reporDNA[8]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[9]),uint8(reporDNA[9]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[10]),uint8(reporDNA[10]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[11]),uint8(reporDNA[11]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[12]),uint8(reporDNA[12]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[13]),uint8(reporDNA[13]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[14]),uint8(reporDNA[14]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[15]),uint8(reporDNA[15]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[16]),uint8(reporDNA[16]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[17]),uint8(reporDNA[17]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[18]),uint8(reporDNA[18]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[19]),uint8(reporDNA[19]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[20]),uint8(reporDNA[20]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[21]),uint8(reporDNA[21]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[22]),uint8(reporDNA[22]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[23]),uint8(reporDNA[23]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[24]),uint8(reporDNA[24]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[25]),uint8(reporDNA[25]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[26]),uint8(reporDNA[26]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[27]),uint8(reporDNA[27]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[28]),uint8(reporDNA[28]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[29]),uint8(reporDNA[29]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[30]),uint8(reporDNA[30]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[31]),uint8(reporDNA[31]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[32]),uint8(reporDNA[32]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[33]),uint8(reporDNA[33]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[34]),uint8(reporDNA[34]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[35]),uint8(reporDNA[35]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[36]),uint8(reporDNA[36]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[37]),uint8(reporDNA[37]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[38]),uint8(reporDNA[38]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[39]),uint8(reporDNA[39]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[40]),uint8(reporDNA[40]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[41]),uint8(reporDNA[41]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[42]),uint8(reporDNA[42]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[43]),uint8(reporDNA[43]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[44]),uint8(reporDNA[44]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[45]),uint8(reporDNA[45]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[46]),uint8(reporDNA[46]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[47]),uint8(reporDNA[47]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[48]),uint8(reporDNA[48]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[49]),uint8(reporDNA[49]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[50]),uint8(reporDNA[50]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[51]),uint8(reporDNA[51]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[52]),uint8(reporDNA[52]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[53]),uint8(reporDNA[53]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[54]),uint8(reporDNA[54]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[55]),uint8(reporDNA[55]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[56]),uint8(reporDNA[56]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[57]),uint8(reporDNA[57]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[58]),uint8(reporDNA[58]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[59]),uint8(reporDNA[59]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[60]),uint8(reporDNA[60]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[61]),uint8(reporDNA[61]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[62]),uint8(reporDNA[62]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[63]),uint8(reporDNA[63]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[64]),uint8(reporDNA[64]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[65]),uint8(reporDNA[65]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[66]),uint8(reporDNA[66]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[67]),uint8(reporDNA[67]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[68]),uint8(reporDNA[68]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[69]),uint8(reporDNA[69]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[70]),uint8(reporDNA[70]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[71]),uint8(reporDNA[71]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[72]),uint8(reporDNA[72]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[73]),uint8(reporDNA[73]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[74]),uint8(reporDNA[74]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[75]),uint8(reporDNA[75]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[76]),uint8(reporDNA[76]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[77]),uint8(reporDNA[77]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[78]),uint8(reporDNA[78]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[79]),uint8(reporDNA[79]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[80]),uint8(reporDNA[80]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[81]),uint8(reporDNA[81]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[82]),uint8(reporDNA[82]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[83]),uint8(reporDNA[83]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[84]),uint8(reporDNA[84]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[85]),uint8(reporDNA[85]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[86]),uint8(reporDNA[86]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[87]),uint8(reporDNA[87]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[88]),uint8(reporDNA[88]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[89]),uint8(reporDNA[89]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[90]),uint8(reporDNA[90]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[91]),uint8(reporDNA[91]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[92]),uint8(reporDNA[92]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[93]),uint8(reporDNA[93]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[94]),uint8(reporDNA[94]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[95]),uint8(reporDNA[95]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[96]),uint8(reporDNA[96]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[97]),uint8(reporDNA[97]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[98]),uint8(reporDNA[98]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[99]),uint8(reporDNA[99]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[100]),uint8(reporDNA[100]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[101]),uint8(reporDNA[101]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[102]),uint8(reporDNA[102]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[103]),uint8(reporDNA[103]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[104]),uint8(reporDNA[104]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[105]),uint8(reporDNA[105]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[106]),uint8(reporDNA[106]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[107]),uint8(reporDNA[107]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[108]),uint8(reporDNA[108]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[109]),uint8(reporDNA[109]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[110]),uint8(reporDNA[110]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[111]),uint8(reporDNA[111]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[112]),uint8(reporDNA[112]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[113]),uint8(reporDNA[113]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[114]),uint8(reporDNA[114]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[115]),uint8(reporDNA[115]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[116]),uint8(reporDNA[116]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[117]),uint8(reporDNA[117]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[118]),uint8(reporDNA[118]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[119]),uint8(reporDNA[119]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[120]),uint8(reporDNA[120]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[121]),uint8(reporDNA[121]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[122]),uint8(reporDNA[122]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[123]),uint8(reporDNA[123]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[124]),uint8(reporDNA[124]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[125]),uint8(reporDNA[125]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[126]),uint8(reporDNA[126]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[127]),uint8(reporDNA[127]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[128]),uint8(reporDNA[128]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[129]),uint8(reporDNA[129]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[130]),uint8(reporDNA[130]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[131]),uint8(reporDNA[131]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[132]),uint8(reporDNA[132]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[133]),uint8(reporDNA[133]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[134]),uint8(reporDNA[134]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[135]),uint8(reporDNA[135]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[136]),uint8(reporDNA[136]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[137]),uint8(reporDNA[137]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[138]),uint8(reporDNA[138]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[139]),uint8(reporDNA[139]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[140]),uint8(reporDNA[140]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[141]),uint8(reporDNA[141]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[142]),uint8(reporDNA[142]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
		r,p = uint8(ref_dna[143]),uint8(reporDNA[143]);if ( r > p) {x = r - p} else { x = p - r }; if (x > SHORTDISTANCE) {return false;} else{runnintTotal = runnintTotal + uint16(x * x);};
	*/

	if ( uint8(ref_dna[1]) > uint8(reporDNA[  1])) {
		x = uint8(ref_dna[1]) - uint8(reporDNA[  1])
	} else {
		x = uint8(reporDNA[  1]) - uint8(ref_dna[1])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[2]) > uint8(reporDNA[  2])) {
		x = uint8(ref_dna[2]) - uint8(reporDNA[  2])
	} else {
		x = uint8(reporDNA[  2]) - uint8(ref_dna[2])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[3]) > uint8(reporDNA[  3])) {
		x = uint8(ref_dna[3]) - uint8(reporDNA[  3])
	} else {
		x = uint8(reporDNA[  3]) - uint8(ref_dna[3])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[4]) > uint8(reporDNA[  4])) {
		x = uint8(ref_dna[4]) - uint8(reporDNA[  4])
	} else {
		x = uint8(reporDNA[  4]) - uint8(ref_dna[4])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[5]) > uint8(reporDNA[  5])) {
		x = uint8(ref_dna[5]) - uint8(reporDNA[  5])
	} else {
		x = uint8(reporDNA[  5]) - uint8(ref_dna[5])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[6]) > uint8(reporDNA[  6])) {
		x = uint8(ref_dna[6]) - uint8(reporDNA[  6])
	} else {
		x = uint8(reporDNA[  6]) - uint8(ref_dna[6])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[7]) > uint8(reporDNA[  7])) {
		x = uint8(ref_dna[7]) - uint8(reporDNA[  7])
	} else {
		x = uint8(reporDNA[  7]) - uint8(ref_dna[7])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[8]) > uint8(reporDNA[  8])) {
		x = uint8(ref_dna[8]) - uint8(reporDNA[  8])
	} else {
		x = uint8(reporDNA[  8]) - uint8(ref_dna[8])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[9]) > uint8(reporDNA[  9])) {
		x = uint8(ref_dna[9]) - uint8(reporDNA[  9])
	} else {
		x = uint8(reporDNA[  9]) - uint8(ref_dna[9])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[10]) > uint8(reporDNA[ 10])) {
		x = uint8(ref_dna[10]) - uint8(reporDNA[ 10])
	} else {
		x = uint8(reporDNA[ 10]) - uint8(ref_dna[10])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[11]) > uint8(reporDNA[ 11])) {
		x = uint8(ref_dna[11]) - uint8(reporDNA[ 11])
	} else {
		x = uint8(reporDNA[ 11]) - uint8(ref_dna[11])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[12]) > uint8(reporDNA[ 12])) {
		x = uint8(ref_dna[12]) - uint8(reporDNA[ 12])
	} else {
		x = uint8(reporDNA[ 12]) - uint8(ref_dna[12])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[13]) > uint8(reporDNA[ 13])) {
		x = uint8(ref_dna[13]) - uint8(reporDNA[ 13])
	} else {
		x = uint8(reporDNA[ 13]) - uint8(ref_dna[13])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[14]) > uint8(reporDNA[ 14])) {
		x = uint8(ref_dna[14]) - uint8(reporDNA[ 14])
	} else {
		x = uint8(reporDNA[ 14]) - uint8(ref_dna[14])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[15]) > uint8(reporDNA[ 15])) {
		x = uint8(ref_dna[15]) - uint8(reporDNA[ 15])
	} else {
		x = uint8(reporDNA[ 15]) - uint8(ref_dna[15])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[16]) > uint8(reporDNA[ 16])) {
		x = uint8(ref_dna[16]) - uint8(reporDNA[ 16])
	} else {
		x = uint8(reporDNA[ 16]) - uint8(ref_dna[16])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[17]) > uint8(reporDNA[ 17])) {
		x = uint8(ref_dna[17]) - uint8(reporDNA[ 17])
	} else {
		x = uint8(reporDNA[ 17]) - uint8(ref_dna[17])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[18]) > uint8(reporDNA[ 18])) {
		x = uint8(ref_dna[18]) - uint8(reporDNA[ 18])
	} else {
		x = uint8(reporDNA[ 18]) - uint8(ref_dna[18])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[19]) > uint8(reporDNA[ 19])) {
		x = uint8(ref_dna[19]) - uint8(reporDNA[ 19])
	} else {
		x = uint8(reporDNA[ 19]) - uint8(ref_dna[19])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[20]) > uint8(reporDNA[ 20])) {
		x = uint8(ref_dna[20]) - uint8(reporDNA[ 20])
	} else {
		x = uint8(reporDNA[ 20]) - uint8(ref_dna[20])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[21]) > uint8(reporDNA[ 21])) {
		x = uint8(ref_dna[21]) - uint8(reporDNA[ 21])
	} else {
		x = uint8(reporDNA[ 21]) - uint8(ref_dna[21])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[22]) > uint8(reporDNA[ 22])) {
		x = uint8(ref_dna[22]) - uint8(reporDNA[ 22])
	} else {
		x = uint8(reporDNA[ 22]) - uint8(ref_dna[22])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[23]) > uint8(reporDNA[ 23])) {
		x = uint8(ref_dna[23]) - uint8(reporDNA[ 23])
	} else {
		x = uint8(reporDNA[ 23]) - uint8(ref_dna[23])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[24]) > uint8(reporDNA[ 24])) {
		x = uint8(ref_dna[24]) - uint8(reporDNA[ 24])
	} else {
		x = uint8(reporDNA[ 24]) - uint8(ref_dna[24])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[25]) > uint8(reporDNA[ 25])) {
		x = uint8(ref_dna[25]) - uint8(reporDNA[ 25])
	} else {
		x = uint8(reporDNA[ 25]) - uint8(ref_dna[25])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[26]) > uint8(reporDNA[ 26])) {
		x = uint8(ref_dna[26]) - uint8(reporDNA[ 26])
	} else {
		x = uint8(reporDNA[ 26]) - uint8(ref_dna[26])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[27]) > uint8(reporDNA[ 27])) {
		x = uint8(ref_dna[27]) - uint8(reporDNA[ 27])
	} else {
		x = uint8(reporDNA[ 27]) - uint8(ref_dna[27])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[28]) > uint8(reporDNA[ 28])) {
		x = uint8(ref_dna[28]) - uint8(reporDNA[ 28])
	} else {
		x = uint8(reporDNA[ 28]) - uint8(ref_dna[28])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[29]) > uint8(reporDNA[ 29])) {
		x = uint8(ref_dna[29]) - uint8(reporDNA[ 29])
	} else {
		x = uint8(reporDNA[ 29]) - uint8(ref_dna[29])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[30]) > uint8(reporDNA[ 30])) {
		x = uint8(ref_dna[30]) - uint8(reporDNA[ 30])
	} else {
		x = uint8(reporDNA[ 30]) - uint8(ref_dna[30])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[31]) > uint8(reporDNA[ 31])) {
		x = uint8(ref_dna[31]) - uint8(reporDNA[ 31])
	} else {
		x = uint8(reporDNA[ 31]) - uint8(ref_dna[31])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[32]) > uint8(reporDNA[ 32])) {
		x = uint8(ref_dna[32]) - uint8(reporDNA[ 32])
	} else {
		x = uint8(reporDNA[ 32]) - uint8(ref_dna[32])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[33]) > uint8(reporDNA[ 33])) {
		x = uint8(ref_dna[33]) - uint8(reporDNA[ 33])
	} else {
		x = uint8(reporDNA[ 33]) - uint8(ref_dna[33])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[34]) > uint8(reporDNA[ 34])) {
		x = uint8(ref_dna[34]) - uint8(reporDNA[ 34])
	} else {
		x = uint8(reporDNA[ 34]) - uint8(ref_dna[34])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[35]) > uint8(reporDNA[ 35])) {
		x = uint8(ref_dna[35]) - uint8(reporDNA[ 35])
	} else {
		x = uint8(reporDNA[ 35]) - uint8(ref_dna[35])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[36]) > uint8(reporDNA[ 36])) {
		x = uint8(ref_dna[36]) - uint8(reporDNA[ 36])
	} else {
		x = uint8(reporDNA[ 36]) - uint8(ref_dna[36])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[37]) > uint8(reporDNA[ 37])) {
		x = uint8(ref_dna[37]) - uint8(reporDNA[ 37])
	} else {
		x = uint8(reporDNA[ 37]) - uint8(ref_dna[37])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[38]) > uint8(reporDNA[ 38])) {
		x = uint8(ref_dna[38]) - uint8(reporDNA[ 38])
	} else {
		x = uint8(reporDNA[ 38]) - uint8(ref_dna[38])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[39]) > uint8(reporDNA[ 39])) {
		x = uint8(ref_dna[39]) - uint8(reporDNA[ 39])
	} else {
		x = uint8(reporDNA[ 39]) - uint8(ref_dna[39])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[40]) > uint8(reporDNA[ 40])) {
		x = uint8(ref_dna[40]) - uint8(reporDNA[ 40])
	} else {
		x = uint8(reporDNA[ 40]) - uint8(ref_dna[40])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[41]) > uint8(reporDNA[ 41])) {
		x = uint8(ref_dna[41]) - uint8(reporDNA[ 41])
	} else {
		x = uint8(reporDNA[ 41]) - uint8(ref_dna[41])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[42]) > uint8(reporDNA[ 42])) {
		x = uint8(ref_dna[42]) - uint8(reporDNA[ 42])
	} else {
		x = uint8(reporDNA[ 42]) - uint8(ref_dna[42])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[43]) > uint8(reporDNA[ 43])) {
		x = uint8(ref_dna[43]) - uint8(reporDNA[ 43])
	} else {
		x = uint8(reporDNA[ 43]) - uint8(ref_dna[43])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[44]) > uint8(reporDNA[ 44])) {
		x = uint8(ref_dna[44]) - uint8(reporDNA[ 44])
	} else {
		x = uint8(reporDNA[ 44]) - uint8(ref_dna[44])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[45]) > uint8(reporDNA[ 45])) {
		x = uint8(ref_dna[45]) - uint8(reporDNA[ 45])
	} else {
		x = uint8(reporDNA[ 45]) - uint8(ref_dna[45])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[46]) > uint8(reporDNA[ 46])) {
		x = uint8(ref_dna[46]) - uint8(reporDNA[ 46])
	} else {
		x = uint8(reporDNA[ 46]) - uint8(ref_dna[46])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[47]) > uint8(reporDNA[ 47])) {
		x = uint8(ref_dna[47]) - uint8(reporDNA[ 47])
	} else {
		x = uint8(reporDNA[ 47]) - uint8(ref_dna[47])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[48]) > uint8(reporDNA[ 48])) {
		x = uint8(ref_dna[48]) - uint8(reporDNA[ 48])
	} else {
		x = uint8(reporDNA[ 48]) - uint8(ref_dna[48])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[49]) > uint8(reporDNA[ 49])) {
		x = uint8(ref_dna[49]) - uint8(reporDNA[ 49])
	} else {
		x = uint8(reporDNA[ 49]) - uint8(ref_dna[49])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[50]) > uint8(reporDNA[ 50])) {
		x = uint8(ref_dna[50]) - uint8(reporDNA[ 50])
	} else {
		x = uint8(reporDNA[ 50]) - uint8(ref_dna[50])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[51]) > uint8(reporDNA[ 51])) {
		x = uint8(ref_dna[51]) - uint8(reporDNA[ 51])
	} else {
		x = uint8(reporDNA[ 51]) - uint8(ref_dna[51])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[52]) > uint8(reporDNA[ 52])) {
		x = uint8(ref_dna[52]) - uint8(reporDNA[ 52])
	} else {
		x = uint8(reporDNA[ 52]) - uint8(ref_dna[52])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[53]) > uint8(reporDNA[ 53])) {
		x = uint8(ref_dna[53]) - uint8(reporDNA[ 53])
	} else {
		x = uint8(reporDNA[ 53]) - uint8(ref_dna[53])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[54]) > uint8(reporDNA[ 54])) {
		x = uint8(ref_dna[54]) - uint8(reporDNA[ 54])
	} else {
		x = uint8(reporDNA[ 54]) - uint8(ref_dna[54])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[55]) > uint8(reporDNA[ 55])) {
		x = uint8(ref_dna[55]) - uint8(reporDNA[ 55])
	} else {
		x = uint8(reporDNA[ 55]) - uint8(ref_dna[55])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[56]) > uint8(reporDNA[ 56])) {
		x = uint8(ref_dna[56]) - uint8(reporDNA[ 56])
	} else {
		x = uint8(reporDNA[ 56]) - uint8(ref_dna[56])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[57]) > uint8(reporDNA[ 57])) {
		x = uint8(ref_dna[57]) - uint8(reporDNA[ 57])
	} else {
		x = uint8(reporDNA[ 57]) - uint8(ref_dna[57])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[58]) > uint8(reporDNA[ 58])) {
		x = uint8(ref_dna[58]) - uint8(reporDNA[ 58])
	} else {
		x = uint8(reporDNA[ 58]) - uint8(ref_dna[58])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[59]) > uint8(reporDNA[ 59])) {
		x = uint8(ref_dna[59]) - uint8(reporDNA[ 59])
	} else {
		x = uint8(reporDNA[ 59]) - uint8(ref_dna[59])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[60]) > uint8(reporDNA[ 60])) {
		x = uint8(ref_dna[60]) - uint8(reporDNA[ 60])
	} else {
		x = uint8(reporDNA[ 60]) - uint8(ref_dna[60])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[61]) > uint8(reporDNA[ 61])) {
		x = uint8(ref_dna[61]) - uint8(reporDNA[ 61])
	} else {
		x = uint8(reporDNA[ 61]) - uint8(ref_dna[61])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[62]) > uint8(reporDNA[ 62])) {
		x = uint8(ref_dna[62]) - uint8(reporDNA[ 62])
	} else {
		x = uint8(reporDNA[ 62]) - uint8(ref_dna[62])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[63]) > uint8(reporDNA[ 63])) {
		x = uint8(ref_dna[63]) - uint8(reporDNA[ 63])
	} else {
		x = uint8(reporDNA[ 63]) - uint8(ref_dna[63])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[64]) > uint8(reporDNA[ 64])) {
		x = uint8(ref_dna[64]) - uint8(reporDNA[ 64])
	} else {
		x = uint8(reporDNA[ 64]) - uint8(ref_dna[64])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[65]) > uint8(reporDNA[ 65])) {
		x = uint8(ref_dna[65]) - uint8(reporDNA[ 65])
	} else {
		x = uint8(reporDNA[ 65]) - uint8(ref_dna[65])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[66]) > uint8(reporDNA[ 66])) {
		x = uint8(ref_dna[66]) - uint8(reporDNA[ 66])
	} else {
		x = uint8(reporDNA[ 66]) - uint8(ref_dna[66])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[67]) > uint8(reporDNA[ 67])) {
		x = uint8(ref_dna[67]) - uint8(reporDNA[ 67])
	} else {
		x = uint8(reporDNA[ 67]) - uint8(ref_dna[67])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[68]) > uint8(reporDNA[ 68])) {
		x = uint8(ref_dna[68]) - uint8(reporDNA[ 68])
	} else {
		x = uint8(reporDNA[ 68]) - uint8(ref_dna[68])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[69]) > uint8(reporDNA[ 69])) {
		x = uint8(ref_dna[69]) - uint8(reporDNA[ 69])
	} else {
		x = uint8(reporDNA[ 69]) - uint8(ref_dna[69])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[70]) > uint8(reporDNA[ 70])) {
		x = uint8(ref_dna[70]) - uint8(reporDNA[ 70])
	} else {
		x = uint8(reporDNA[ 70]) - uint8(ref_dna[70])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[71]) > uint8(reporDNA[ 71])) {
		x = uint8(ref_dna[71]) - uint8(reporDNA[ 71])
	} else {
		x = uint8(reporDNA[ 71]) - uint8(ref_dna[71])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[72]) > uint8(reporDNA[ 72])) {
		x = uint8(ref_dna[72]) - uint8(reporDNA[ 72])
	} else {
		x = uint8(reporDNA[ 72]) - uint8(ref_dna[72])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[73]) > uint8(reporDNA[ 73])) {
		x = uint8(ref_dna[73]) - uint8(reporDNA[ 73])
	} else {
		x = uint8(reporDNA[ 73]) - uint8(ref_dna[73])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[74]) > uint8(reporDNA[ 74])) {
		x = uint8(ref_dna[74]) - uint8(reporDNA[ 74])
	} else {
		x = uint8(reporDNA[ 74]) - uint8(ref_dna[74])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[75]) > uint8(reporDNA[ 75])) {
		x = uint8(ref_dna[75]) - uint8(reporDNA[ 75])
	} else {
		x = uint8(reporDNA[ 75]) - uint8(ref_dna[75])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[76]) > uint8(reporDNA[ 76])) {
		x = uint8(ref_dna[76]) - uint8(reporDNA[ 76])
	} else {
		x = uint8(reporDNA[ 76]) - uint8(ref_dna[76])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[77]) > uint8(reporDNA[ 77])) {
		x = uint8(ref_dna[77]) - uint8(reporDNA[ 77])
	} else {
		x = uint8(reporDNA[ 77]) - uint8(ref_dna[77])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[78]) > uint8(reporDNA[ 78])) {
		x = uint8(ref_dna[78]) - uint8(reporDNA[ 78])
	} else {
		x = uint8(reporDNA[ 78]) - uint8(ref_dna[78])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[79]) > uint8(reporDNA[ 79])) {
		x = uint8(ref_dna[79]) - uint8(reporDNA[ 79])
	} else {
		x = uint8(reporDNA[ 79]) - uint8(ref_dna[79])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[80]) > uint8(reporDNA[ 80])) {
		x = uint8(ref_dna[80]) - uint8(reporDNA[ 80])
	} else {
		x = uint8(reporDNA[ 80]) - uint8(ref_dna[80])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[81]) > uint8(reporDNA[ 81])) {
		x = uint8(ref_dna[81]) - uint8(reporDNA[ 81])
	} else {
		x = uint8(reporDNA[ 81]) - uint8(ref_dna[81])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[82]) > uint8(reporDNA[ 82])) {
		x = uint8(ref_dna[82]) - uint8(reporDNA[ 82])
	} else {
		x = uint8(reporDNA[ 82]) - uint8(ref_dna[82])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[83]) > uint8(reporDNA[ 83])) {
		x = uint8(ref_dna[83]) - uint8(reporDNA[ 83])
	} else {
		x = uint8(reporDNA[ 83]) - uint8(ref_dna[83])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[84]) > uint8(reporDNA[ 84])) {
		x = uint8(ref_dna[84]) - uint8(reporDNA[ 84])
	} else {
		x = uint8(reporDNA[ 84]) - uint8(ref_dna[84])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[85]) > uint8(reporDNA[ 85])) {
		x = uint8(ref_dna[85]) - uint8(reporDNA[ 85])
	} else {
		x = uint8(reporDNA[ 85]) - uint8(ref_dna[85])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[86]) > uint8(reporDNA[ 86])) {
		x = uint8(ref_dna[86]) - uint8(reporDNA[ 86])
	} else {
		x = uint8(reporDNA[ 86]) - uint8(ref_dna[86])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[87]) > uint8(reporDNA[ 87])) {
		x = uint8(ref_dna[87]) - uint8(reporDNA[ 87])
	} else {
		x = uint8(reporDNA[ 87]) - uint8(ref_dna[87])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[88]) > uint8(reporDNA[ 88])) {
		x = uint8(ref_dna[88]) - uint8(reporDNA[ 88])
	} else {
		x = uint8(reporDNA[ 88]) - uint8(ref_dna[88])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[89]) > uint8(reporDNA[ 89])) {
		x = uint8(ref_dna[89]) - uint8(reporDNA[ 89])
	} else {
		x = uint8(reporDNA[ 89]) - uint8(ref_dna[89])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[90]) > uint8(reporDNA[ 90])) {
		x = uint8(ref_dna[90]) - uint8(reporDNA[ 90])
	} else {
		x = uint8(reporDNA[ 90]) - uint8(ref_dna[90])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[91]) > uint8(reporDNA[ 91])) {
		x = uint8(ref_dna[91]) - uint8(reporDNA[ 91])
	} else {
		x = uint8(reporDNA[ 91]) - uint8(ref_dna[91])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[92]) > uint8(reporDNA[ 92])) {
		x = uint8(ref_dna[92]) - uint8(reporDNA[ 92])
	} else {
		x = uint8(reporDNA[ 92]) - uint8(ref_dna[92])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[93]) > uint8(reporDNA[ 93])) {
		x = uint8(ref_dna[93]) - uint8(reporDNA[ 93])
	} else {
		x = uint8(reporDNA[ 93]) - uint8(ref_dna[93])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[94]) > uint8(reporDNA[ 94])) {
		x = uint8(ref_dna[94]) - uint8(reporDNA[ 94])
	} else {
		x = uint8(reporDNA[ 94]) - uint8(ref_dna[94])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[95]) > uint8(reporDNA[ 95])) {
		x = uint8(ref_dna[95]) - uint8(reporDNA[ 95])
	} else {
		x = uint8(reporDNA[ 95]) - uint8(ref_dna[95])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[96]) > uint8(reporDNA[ 96])) {
		x = uint8(ref_dna[96]) - uint8(reporDNA[ 96])
	} else {
		x = uint8(reporDNA[ 96]) - uint8(ref_dna[96])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[97]) > uint8(reporDNA[ 97])) {
		x = uint8(ref_dna[97]) - uint8(reporDNA[ 97])
	} else {
		x = uint8(reporDNA[ 97]) - uint8(ref_dna[97])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[98]) > uint8(reporDNA[ 98])) {
		x = uint8(ref_dna[98]) - uint8(reporDNA[ 98])
	} else {
		x = uint8(reporDNA[ 98]) - uint8(ref_dna[98])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if (  uint8(ref_dna[99]) > uint8(reporDNA[ 99])) {
		x = uint8(ref_dna[99]) - uint8(reporDNA[ 99])
	} else {
		x = uint8(reporDNA[ 99]) - uint8(ref_dna[99])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[100]) > uint8(reporDNA[100])) {
		x = uint8(ref_dna[100]) - uint8(reporDNA[100])
	} else {
		x = uint8(reporDNA[100]) - uint8(ref_dna[100])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[101]) > uint8(reporDNA[101])) {
		x = uint8(ref_dna[101]) - uint8(reporDNA[101])
	} else {
		x = uint8(reporDNA[101]) - uint8(ref_dna[101])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[102]) > uint8(reporDNA[102])) {
		x = uint8(ref_dna[102]) - uint8(reporDNA[102])
	} else {
		x = uint8(reporDNA[102]) - uint8(ref_dna[102])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[103]) > uint8(reporDNA[103])) {
		x = uint8(ref_dna[103]) - uint8(reporDNA[103])
	} else {
		x = uint8(reporDNA[103]) - uint8(ref_dna[103])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[104]) > uint8(reporDNA[104])) {
		x = uint8(ref_dna[104]) - uint8(reporDNA[104])
	} else {
		x = uint8(reporDNA[104]) - uint8(ref_dna[104])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[105]) > uint8(reporDNA[105])) {
		x = uint8(ref_dna[105]) - uint8(reporDNA[105])
	} else {
		x = uint8(reporDNA[105]) - uint8(ref_dna[105])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[106]) > uint8(reporDNA[106])) {
		x = uint8(ref_dna[106]) - uint8(reporDNA[106])
	} else {
		x = uint8(reporDNA[106]) - uint8(ref_dna[106])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[107]) > uint8(reporDNA[107])) {
		x = uint8(ref_dna[107]) - uint8(reporDNA[107])
	} else {
		x = uint8(reporDNA[107]) - uint8(ref_dna[107])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[108]) > uint8(reporDNA[108])) {
		x = uint8(ref_dna[108]) - uint8(reporDNA[108])
	} else {
		x = uint8(reporDNA[108]) - uint8(ref_dna[108])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[109]) > uint8(reporDNA[109])) {
		x = uint8(ref_dna[109]) - uint8(reporDNA[109])
	} else {
		x = uint8(reporDNA[109]) - uint8(ref_dna[109])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[110]) > uint8(reporDNA[110])) {
		x = uint8(ref_dna[110]) - uint8(reporDNA[110])
	} else {
		x = uint8(reporDNA[110]) - uint8(ref_dna[110])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[111]) > uint8(reporDNA[111])) {
		x = uint8(ref_dna[111]) - uint8(reporDNA[111])
	} else {
		x = uint8(reporDNA[111]) - uint8(ref_dna[111])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[112]) > uint8(reporDNA[112])) {
		x = uint8(ref_dna[112]) - uint8(reporDNA[112])
	} else {
		x = uint8(reporDNA[112]) - uint8(ref_dna[112])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[113]) > uint8(reporDNA[113])) {
		x = uint8(ref_dna[113]) - uint8(reporDNA[113])
	} else {
		x = uint8(reporDNA[113]) - uint8(ref_dna[113])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[114]) > uint8(reporDNA[114])) {
		x = uint8(ref_dna[114]) - uint8(reporDNA[114])
	} else {
		x = uint8(reporDNA[114]) - uint8(ref_dna[114])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[115]) > uint8(reporDNA[115])) {
		x = uint8(ref_dna[115]) - uint8(reporDNA[115])
	} else {
		x = uint8(reporDNA[115]) - uint8(ref_dna[115])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[116]) > uint8(reporDNA[116])) {
		x = uint8(ref_dna[116]) - uint8(reporDNA[116])
	} else {
		x = uint8(reporDNA[116]) - uint8(ref_dna[116])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[117]) > uint8(reporDNA[117])) {
		x = uint8(ref_dna[117]) - uint8(reporDNA[117])
	} else {
		x = uint8(reporDNA[117]) - uint8(ref_dna[117])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[118]) > uint8(reporDNA[118])) {
		x = uint8(ref_dna[118]) - uint8(reporDNA[118])
	} else {
		x = uint8(reporDNA[118]) - uint8(ref_dna[118])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[119]) > uint8(reporDNA[119])) {
		x = uint8(ref_dna[119]) - uint8(reporDNA[119])
	} else {
		x = uint8(reporDNA[119]) - uint8(ref_dna[119])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[120]) > uint8(reporDNA[120])) {
		x = uint8(ref_dna[120]) - uint8(reporDNA[120])
	} else {
		x = uint8(reporDNA[120]) - uint8(ref_dna[120])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[121]) > uint8(reporDNA[121])) {
		x = uint8(ref_dna[121]) - uint8(reporDNA[121])
	} else {
		x = uint8(reporDNA[121]) - uint8(ref_dna[121])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[122]) > uint8(reporDNA[122])) {
		x = uint8(ref_dna[122]) - uint8(reporDNA[122])
	} else {
		x = uint8(reporDNA[122]) - uint8(ref_dna[122])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[123]) > uint8(reporDNA[123])) {
		x = uint8(ref_dna[123]) - uint8(reporDNA[123])
	} else {
		x = uint8(reporDNA[123]) - uint8(ref_dna[123])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[124]) > uint8(reporDNA[124])) {
		x = uint8(ref_dna[124]) - uint8(reporDNA[124])
	} else {
		x = uint8(reporDNA[124]) - uint8(ref_dna[124])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[125]) > uint8(reporDNA[125])) {
		x = uint8(ref_dna[125]) - uint8(reporDNA[125])
	} else {
		x = uint8(reporDNA[125]) - uint8(ref_dna[125])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[126]) > uint8(reporDNA[126])) {
		x = uint8(ref_dna[126]) - uint8(reporDNA[126])
	} else {
		x = uint8(reporDNA[126]) - uint8(ref_dna[126])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[127]) > uint8(reporDNA[127])) {
		x = uint8(ref_dna[127]) - uint8(reporDNA[127])
	} else {
		x = uint8(reporDNA[127]) - uint8(ref_dna[127])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[128]) > uint8(reporDNA[128])) {
		x = uint8(ref_dna[128]) - uint8(reporDNA[128])
	} else {
		x = uint8(reporDNA[128]) - uint8(ref_dna[128])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[129]) > uint8(reporDNA[129])) {
		x = uint8(ref_dna[129]) - uint8(reporDNA[129])
	} else {
		x = uint8(reporDNA[129]) - uint8(ref_dna[129])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[130]) > uint8(reporDNA[130])) {
		x = uint8(ref_dna[130]) - uint8(reporDNA[130])
	} else {
		x = uint8(reporDNA[130]) - uint8(ref_dna[130])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[131]) > uint8(reporDNA[131])) {
		x = uint8(ref_dna[131]) - uint8(reporDNA[131])
	} else {
		x = uint8(reporDNA[131]) - uint8(ref_dna[131])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[132]) > uint8(reporDNA[132])) {
		x = uint8(ref_dna[132]) - uint8(reporDNA[132])
	} else {
		x = uint8(reporDNA[132]) - uint8(ref_dna[132])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[133]) > uint8(reporDNA[133])) {
		x = uint8(ref_dna[133]) - uint8(reporDNA[133])
	} else {
		x = uint8(reporDNA[133]) - uint8(ref_dna[133])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[134]) > uint8(reporDNA[134])) {
		x = uint8(ref_dna[134]) - uint8(reporDNA[134])
	} else {
		x = uint8(reporDNA[134]) - uint8(ref_dna[134])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[135]) > uint8(reporDNA[135])) {
		x = uint8(ref_dna[135]) - uint8(reporDNA[135])
	} else {
		x = uint8(reporDNA[135]) - uint8(ref_dna[135])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[136]) > uint8(reporDNA[136])) {
		x = uint8(ref_dna[136]) - uint8(reporDNA[136])
	} else {
		x = uint8(reporDNA[136]) - uint8(ref_dna[136])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[137]) > uint8(reporDNA[137])) {
		x = uint8(ref_dna[137]) - uint8(reporDNA[137])
	} else {
		x = uint8(reporDNA[137]) - uint8(ref_dna[137])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[138]) > uint8(reporDNA[138])) {
		x = uint8(ref_dna[138]) - uint8(reporDNA[138])
	} else {
		x = uint8(reporDNA[138]) - uint8(ref_dna[138])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[139]) > uint8(reporDNA[139])) {
		x = uint8(ref_dna[139]) - uint8(reporDNA[139])
	} else {
		x = uint8(reporDNA[139]) - uint8(ref_dna[139])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[140]) > uint8(reporDNA[140])) {
		x = uint8(ref_dna[140]) - uint8(reporDNA[140])
	} else {
		x = uint8(reporDNA[140]) - uint8(ref_dna[140])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[141]) > uint8(reporDNA[141])) {
		x = uint8(ref_dna[141]) - uint8(reporDNA[141])
	} else {
		x = uint8(reporDNA[141]) - uint8(ref_dna[141])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[142]) > uint8(reporDNA[142])) {
		x = uint8(ref_dna[142]) - uint8(reporDNA[142])
	} else {
		x = uint8(reporDNA[142]) - uint8(ref_dna[142])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };
	if ( uint8(ref_dna[143]) > uint8(reporDNA[143])) {
		x = uint8(ref_dna[143]) - uint8(reporDNA[143])
	} else {
		x = uint8(reporDNA[143]) - uint8(ref_dna[143])
	}; if (x > SHORTDISTANCE) {
		return MAXDISTANCE; } else {
		runnintTotal = runnintTotal + uint16(x * x); };

	/*	if (runnintTotal > 40000) {
			// || runnintTotal>40000


			return 40001
		} else {
			//fmt.Println("Match")
			return true
		}

	*/
	return int(runnintTotal)

	//	runnintTotal=(runnintTotal+x)

}

func pdnaCompare(ref_dna []byte, reporDNA []byte) bool {
	//var maxTotal = 2000
	var runnintTotal uint16
	var x uint8
	var r uint8
	var p uint8
	for ii := 0; ii < PDNALEN; ii++ {
		r = uint8(ref_dna[ii])
		p = uint8(reporDNA[ii])
		if ( r > p) {
			x = r - p
		} else {
			x = p - r
		}

		if (x > SHORTDISTANCE) {
			//fmt.Print("No Match ")
			runnintTotal = MAXDISTANCE
			break
		}
		runnintTotal = runnintTotal + uint16(x * x)
	}
	if (runnintTotal >= MAXDISTANCE) {
		// || runnintTotal>40000


		return false
	} else {
		//fmt.Println("Match")
		return true
	}


	//	runnintTotal=(runnintTotal+x)

}
func ComparePDNA(dat []byte, reportDNA []byte, ch chan int, threadNum int) {
	var datlen int
	var replen int
	datlen = len(dat) / PDNALEN
	replen = len(reportDNA) / PDNALEN
	var foundcount int

	//t := time.Now()
	//fmt.Println(t)
	for i := 0; i < datlen; i++ {
		startIdx := i * PDNALEN
		ref_dna := dat[startIdx:startIdx + PDNALEN]
		for j := 0; j < replen; j++ {
			startIdx2 := j * PDNALEN
			report_dna := reportDNA[startIdx2:startIdx2 + PDNALEN]
			//_ = pdnaCompare(ref_dna, report_dna)
			distance := pdnaCompareStatic(ref_dna, report_dna)
			if (distance < MAXDISTANCE) {
				foundcount++
			}
		}
		/*
		for j := 1; j <= PDNALEN; j++ {
			fmt.Print(dat[(i * PDNALEN) - j])

		}*/
		//fmt.Println("")
		//ComparePDNA(refDNA,reportDNA)
	}
	//elapsed := time.Since(t)
	// fmt.Println("ThreadNumber:",threadNum,"Refset per Thread:",datlen, " Report Set Size:" ,replen,"Compare took %s", elapsed)
	ch <- foundcount
}
func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func DoPDNAWORK(refenceSet []byte, reportSet[]byte) {

	fmt.Println(MaxParallelism(), "Number CPU")
	cpu := MaxParallelism() - 2
	//chunk := int(len(refenceSet) /cpu)
	chunk := int(len(refenceSet) / cpu) / PDNALEN
	fmt.Println("Chunk Size", chunk, chunk / PDNALEN, chunk)
	reportSlice := reportSet[0:]
	ch := make(chan int, cpu)
	for i := 0; i < cpu; i++ {

		var s []byte

		if (i == (cpu - 1)) {
			s = refenceSet[i * chunk * PDNALEN:]
			workSize := (len(refenceSet) / PDNALEN) - (i * chunk)
			fmt.Println("Thread: ", i + 1, "-", i * chunk, "start-end", len(refenceSet) / PDNALEN, "WorkSize:", workSize)
		} else {
			s = refenceSet[i * chunk * PDNALEN:(chunk * i + chunk - 1) * PDNALEN]
			workSize := (chunk * i + (chunk - 1 )) - (i * chunk)
			fmt.Println("Thread: ", i + 1, "-", i * chunk, "start-end", chunk * i + (chunk - 1 ), "WorkSize:", workSize)

		}

		//pass in slices don't pass arrays
		go ComparePDNA(s, reportSlice, ch, i)
		//time.Sleep(1000 * time.Millisecond)
	}


	//wait for each thread to finish
	for i := 0; i < cpu; i++ {
		foundcount := <-ch
		fmt.Println("Matches Found:", foundcount)
	}
	fmt.Println("PDNA compare Complete")


}
