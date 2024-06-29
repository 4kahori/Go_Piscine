/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   recursivepower.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/29 14:18:07 by kahori            #+#    #+#             */
/*   Updated: 2024/06/29 14:19:41 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

//for 使えない！
// nb の power乗を返す
package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}
