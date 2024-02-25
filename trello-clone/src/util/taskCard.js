import axios from "axios"

const taskCardUtil = () => {
	const registCard = async(title, sort_no) => {
		const res = await axios.post(
			`${process.env.REACT_APP_API_URL}/tasks/card`,
			{
				title: title,
				sort_no: sort_no,
			},
		)
		return res
	}

	const updateCard = async(id, title, sort_no) => {
		const res = await axios.put(
			`${process.env.REACT_APP_API_URL}/tasks/card/${id}`,
			{
				title: title,
				sort_no: sort_no,
			},
		)
		return res
	}

	const deleteCard = async(id) => {
		const res = await axios.delete(
			`${process.env.REACT_APP_API_URL}/tasks/card/${id}`,
		)
		return res
	}
	// const login = async(email, password) => {
	// 	const res = await axios.post(
	// 		`${process.env.REACT_APP_API_URL}/login`,
	// 		{
	// 			email:email,
	// 			password:password,
	// 		},
	// 		// {withCredentials: true,}
	// 	)
	// 	return res
	// }
	// const logout = async() => {
	// 	const res = await axios.post(
	// 		`${process.env.REACT_APP_API_URL}/logout`,
	// 		// {withCredentials: true}
	// 	)
	// 	return res
	// }
	return {registCard, updateCard, deleteCard}
}

export default taskCardUtil