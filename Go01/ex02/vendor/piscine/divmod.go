/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   divmod.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/27 14:08:21 by kahori            #+#    #+#             */
/*   Updated: 2024/06/27 14:08:22 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}
