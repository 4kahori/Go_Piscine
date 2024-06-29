/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sqrt.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/29 15:16:12 by kahori            #+#    #+#             */
/*   Updated: 2024/06/29 15:18:40 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
