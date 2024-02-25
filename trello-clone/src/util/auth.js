import axios from "axios"

const auth = () => {
	const signup = async(email, password) => {
		const res = await axios.post(
			`${process.env.REACT_APP_API_URL}/signup`,
			{
				email: email,
				password: password,
			},
		)
		return res
	}
	const login = async(email, password) => {
		const res = await axios.post(
			`${process.env.REACT_APP_API_URL}/login`,
			{
				email:email,
				password:password,
			},
			{withCredentials: true}
		)
		return res
	}
	const logout = async() => {
		const res = await axios.post(
			`${process.env.REACT_APP_API_URL}/logout`,
			{withCredentials: true}
		)
		return res
	}
	return {signup, login, logout}
}

export default auth