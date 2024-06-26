/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   atoi.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/27 16:31:18 by kahori            #+#    #+#             */
/*   Updated: 2024/06/27 16:33:47 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func Atoi(s string) int {
	n := 0
	sign := 1
	for i, r := range s {
		if i == 0 && (r == '-' || r == '+') {
			if r == '-' {
				sign = -1
			}
			continue
		}
		if r < '0' || r > '9' {
			return 0
		}
		n = n * 10 + int(r - '0')
	}
	return n * sign
}
