/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   findnextprime.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/29 16:37:09 by kahori            #+#    #+#             */
/*   Updated: 2024/06/29 16:44:09 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func FindNextPrime(nb int) int {
    if nb <= 1 {
        return 2
    }
    for !IsPrime(nb) {
        nb++
    }
    return nb
}