package datasize

import "fmt"
import "github.com/kormoc/unit"
import "math"
import "strings"

type Datasize        float64
type DatasizeSIBit   Datasize
type DatasizeSIByte  Datasize
type DatasizeIECBit  Datasize
type DatasizeIECByte Datasize

var outputStringMaxPercision int = 3
var outputStringMaxLevels    int = 1

const Bit       Datasize = 1
const Kilobit   Datasize = Bit       * 1000
const Megabit   Datasize = Kilobit   * 1000
const Gigabit   Datasize = Megabit   * 1000
const Terabit   Datasize = Gigabit   * 1000
const Petabit   Datasize = Terabit   * 1000
const Exabit    Datasize = Petabit   * 1000
const Zettabit  Datasize = Exabit    * 1000
const Yottabit  Datasize = Zettabit  * 1000

var SIBits = [...]unit.Unit {
    {Name: "yottabit", Size: float64(Yottabit), Suffix: "Ybit"},
    {Name: "zettabit", Size: float64(Zettabit), Suffix: "Zbit"},
    {Name:   "exabit", Size: float64(  Exabit), Suffix: "Ebit"},
    {Name:  "petabit", Size: float64( Petabit), Suffix: "Pbit"},
    {Name:  "terabit", Size: float64( Terabit), Suffix: "Tbit"},
    {Name:  "gigabit", Size: float64( Gigabit), Suffix: "Gbit"},
    {Name:  "megabit", Size: float64( Megabit), Suffix: "Mbit"},
    {Name:  "kilobit", Size: float64( Kilobit), Suffix: "kbit"},
    {Name:      "bit", Size: float64(     Bit), Suffix: "bit"},
}

const Kibibit   Datasize = Bit       * 1024
const Mebibit   Datasize = Kibibit   * 1024
const Gibibit   Datasize = Mebibit   * 1024
const Tebibit   Datasize = Gibibit   * 1024
const Pebibit   Datasize = Tebibit   * 1024
const Exbibit   Datasize = Pebibit   * 1024
const Zebibit   Datasize = Exbibit   * 1024
const Yobibit   Datasize = Zebibit   * 1024

var IECBits = [...]unit.Unit {
    {Name: "yobibit", Size: float64(Yobibit), Suffix: "Yibit"},
    {Name: "zebibit", Size: float64(Zebibit), Suffix: "Zibit"},
    {Name: "exbibit", Size: float64(Exbibit), Suffix: "Eibit"},
    {Name: "pebibit", Size: float64(Pebibit), Suffix: "Pibit"},
    {Name: "tebibit", Size: float64(Tebibit), Suffix: "Tibit"},
    {Name: "gibibit", Size: float64(Gibibit), Suffix: "Gibit"},
    {Name: "mebibit", Size: float64(Mebibit), Suffix: "Mibit"},
    {Name: "kibibit", Size: float64(Kibibit), Suffix: "Kibit"},
    {Name:     "bit", Size: float64(    Bit), Suffix: "bit"},
}

const Byte      Datasize = 8
const Kilobyte  Datasize = Byte      * 1000
const Megabyte  Datasize = Kilobyte  * 1000
const Gigabyte  Datasize = Megabyte  * 1000
const Terabyte  Datasize = Gigabyte  * 1000
const Petabyte  Datasize = Terabyte  * 1000
const Exabyte   Datasize = Petabyte  * 1000
const Zettabyte Datasize = Exabyte   * 1000
const Yottabyte Datasize = Zettabyte * 1000

var SIBytes = [...]unit.Unit {
    {Name: "yottabyte", Size: float64(Yottabyte), Suffix: "Y"},
    {Name: "zettabyte", Size: float64(Zettabyte), Suffix: "Z"},
    {Name:   "exabyte", Size: float64(  Exabyte), Suffix: "E"},
    {Name:  "petabyte", Size: float64( Petabyte), Suffix: "P"},
    {Name:  "terabyte", Size: float64( Terabyte), Suffix: "T"},
    {Name:  "gigabyte", Size: float64( Gigabyte), Suffix: "G"},
    {Name:  "megabyte", Size: float64( Megabyte), Suffix: "M"},
    {Name:  "kilobyte", Size: float64( Kilobyte), Suffix: "k"},
    {Name:      "byte", Size: float64(     Byte), Suffix: "b"},
}

const Kibibyte  Datasize = Byte      * 1024
const Mebibyte  Datasize = Kibibyte  * 1024
const Gibibyte  Datasize = Mebibyte  * 1024
const Tebibyte  Datasize = Gibibyte  * 1024
const Pebibyte  Datasize = Tebibyte  * 1024
const Exbibyte  Datasize = Pebibyte  * 1024
const Zebibyte  Datasize = Exbibyte  * 1024
const Yobibyte  Datasize = Zebibyte  * 1024

var IECBytes = [...]unit.Unit {
    {Name: "yobibyte", Size: float64(Yobibyte), Suffix: "Yi"},
    {Name: "zebibyte", Size: float64(Zebibyte), Suffix: "Zi"},
    {Name: "exbibyte", Size: float64(Exbibyte), Suffix: "Ei"},
    {Name: "pebibyte", Size: float64(Pebibyte), Suffix: "Pi"},
    {Name: "tebibyte", Size: float64(Tebibyte), Suffix: "Ti"},
    {Name: "gibibyte", Size: float64(Gibibyte), Suffix: "Gi"},
    {Name: "mebibyte", Size: float64(Mebibyte), Suffix: "Mi"},
    {Name: "kibibyte", Size: float64(Kibibyte), Suffix: "Ki"},
    {Name:     "byte", Size: float64(    Byte), Suffix: "b"},
}

func SetOutputStringMaxPercision(maxPercision int) {
    outputStringMaxPercision = maxPercision

}

func SetOutputStringMaxLevels(maxLevels int) {
    outputStringMaxLevels = maxLevels
}

func buildHumanString(value float64, mapping [9]unit.Unit) (dataSize string) {
    level := 0
    for _, v := range mapping {
        if value >= v.Size {
            level++

            if level == outputStringMaxLevels {
                // Last level to process, display a float
                vals := strings.SplitN(fmt.Sprintf("%f", value / v.Size), ".", 2)
                dataSize += vals[0]
                dataSize += strings.TrimRight("."+vals[1][:outputStringMaxPercision], ".0")
                dataSize += v.Suffix
                break
            } else {
                // Add current level and continue
                valRemainder := math.Mod(value, v.Size)
                valLevel := int((value - valRemainder)/v.Size)
                value = valRemainder

                dataSize += fmt.Sprintf("%d%v", valLevel, v.Suffix)
            }
        }
    }
    return dataSize
}

func (d DatasizeSIBit) String() string {
    return buildHumanString(float64(d), SIBits)
}

func (d DatasizeSIByte) String() string {
    return buildHumanString(float64(d), SIBytes)
}

func (d DatasizeIECBit) String() string {
    return buildHumanString(float64(d), IECBits)
}

func (d DatasizeIECByte) String() string {
    return buildHumanString(float64(d), IECBytes)
}
