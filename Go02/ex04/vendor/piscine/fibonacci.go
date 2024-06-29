/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   fibonacci.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/29 14:25:02 by kahori            #+#    #+#             */
/*   Updated: 2024/06/29 14:36:55 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index < 2 {
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
