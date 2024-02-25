import React from 'react'
import { IoIosLogOut } from "react-icons/io";
import { IconContext } from "react-icons"
import auth from "../../util/auth";
import { useNavigate } from "react-router-dom";

export const Header = () => {
	const { logout } = auth();
	const navigate = useNavigate()
	const doLogout = async() => {
		try {
			await logout();
			navigate("/");
		}catch(e) {
			console.error(e);
		}
	}
	return (
		<div>
			<header className="flex justify-between items-center">
				<div></div>
				<h1 className="text-4xl">Simple Trello</h1>
				<IconContext.Provider value={{ color: '#ffffff', size: '36px' }}>
					<div className="justify-end">
						<IoIosLogOut className="cursor-pointer" onClick={doLogout} />
					</div>
				</IconContext.Provider>
			</header>
		</div>
	)
}
