import axios from "axios"

const taskUtil = () => {
	const getTaskCardAndTask = async() => {
		const res = await axios.get(
			`${process.env.REACT_APP_API_URL}/tasks`,
			{withCredentials: true}
		)
		return res
	}

	const registTask = async(content, task_card_id, sort_no) => {
		const res = await axios.post(
			`${process.env.REACT_APP_API_URL}/tasks/task`,
			{
				content: content,
				task_card_id: task_card_id,
				sort_no: sort_no,
			},
		)
		return res
	}

	const updateTask = async(id, content, sort_no) => {
		const res = await axios.put(
			`${process.env.REACT_APP_API_URL}/tasks/task/${id}`,
			{
				content: content,
				sort_no: sort_no,
			},
		)
		return res
	}

	const deleteTask = async(id, task_card_id, sort_no) => {
		const res = await axios.delete(
			`${process.env.REACT_APP_API_URL}/tasks/task/${id}`,
			{
				task_card_id: task_card_id,
				sort_no: sort_no,
			},
		)
		return res;
	};
	return { getTaskCardAndTask, registTask, updateTask, deleteTask }
}

export default taskUtil