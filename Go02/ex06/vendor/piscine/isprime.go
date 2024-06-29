/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isprime.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/29 15:46:12 by kahori            #+#    #+#             */
/*   Updated: 2024/06/29 16:35:32 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if (nb % i == 0){
			return false
		}
	}
	return true
}
