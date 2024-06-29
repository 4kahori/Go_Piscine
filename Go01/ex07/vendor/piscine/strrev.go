/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   strrev.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/27 16:15:38 by kahori            #+#    #+#             */
/*   Updated: 2024/06/27 16:28:25 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func StrRev(s string) string {
	runes := []rune(s)
	count := len(s) - 1
	for _,r := range s {
		runes[count] = r
		count--
	}
	return string(runes)
}
