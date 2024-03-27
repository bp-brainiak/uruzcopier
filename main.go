/*
Copyright © 2024 Ulises Ruz Puga <ulises.ruz@gmail.com>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; version 2
of the License.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"fmt"
	"uruzcopier/cmd"
)

func main() {
	fmt.Println("uruzcopier v 0.0.1a (c) 2024 Ulises Ruz Puga, released under GPLv2")
	fmt.Println("this software comes with ABSOLUTELY NO WARRANTY")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Println("under certain conditions; see the GNU General Public License v2 for details")
	cmd.Execute()
}
