/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ultimatedivmod.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: kahori <kahori@student.42tokyo.jp>         +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/27 14:07:30 by kahori            #+#    #+#             */
/*   Updated: 2024/06/27 14:33:38 by kahori           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	mod := *a % *b
	
	*a = div
	*b = mod
}
