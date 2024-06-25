/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printreversealphabet.go                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 14:59:44 by kahori            #+#    #+#             */
/*   Updated: 2024/06/25 14:59:46 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import(
	"ft"
)

func PrintReverseAlphabet(){
	for c := 'z'; c >='a'; c--{
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
