/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isnegative.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 14:35:47 by kahori            #+#    #+#             */
/*   Updated: 2024/06/25 14:38:57 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"ft"
)

func IsNegative(nb int) {
	if nb < 0 {
		ft.PrintRune('T')
	} else {
		ft.PrintRune('F')
	}
	ft.PrintRune('\n')
}
