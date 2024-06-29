/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   iterativefactorial.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/29 12:56:09 by kahori            #+#    #+#             */
/*   Updated: 2024/06/29 14:14:23 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}

//nb!