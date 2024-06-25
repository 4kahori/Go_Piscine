/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcomb.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 14:46:44 by kahori            #+#    #+#             */
/*   Updated: 2024/06/25 15:01:15 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"ft"
)

func PrintComb() {
	for a := 0; a <= 7; a++ {
		for b := a + 1; b <= 8; b++ {
			for c := b + 1; c <= 9; c++ {
				ft.PrintRune(rune(a + '0'))
				ft.PrintRune(rune(b + '0'))
				ft.PrintRune(rune(c + '0'))
				if a != 7 {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
	ft.PrintRune('\n')
}
