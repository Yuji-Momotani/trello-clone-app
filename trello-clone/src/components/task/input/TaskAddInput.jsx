import React, { useContext } from 'react'
import { TaskCardContext, TaskListContext } from "../TaskCard";
import taskUtil from "../../../util/task";

export const TaskAddInput = ({
	inputText,
	setInputText,
}) => {

	const [taskList, setTaskList] = useContext(TaskListContext);
	const taskCard = useContext(TaskCardContext);
	const {registTask} = taskUtil();

	const handleSubmit = async(e) => {
		e.preventDefault();

		if (inputText === "") {
			return;
		}
		//カードを追加
		let maxSortNo = -1;
		if (taskList.length === 1) {
			maxSortNo = taskList[0].sort_no;
		} else if (taskList.length > 1) {
			maxSortNo = taskList.reduce((x,y) => x.sort_no > y.sort_no ? x.sort_no : y.sort_no, 0);
		}
		const {data} = await registTask(inputText, taskCard.id ,maxSortNo+1);
		console.log(data);
		setTaskList(prev => [
			...prev, 
			{
				id: data.id,
				contetn: data.contetn,
				sort_no: data.sort_no,
				task_card_id: data.task_card_id,
			},
		]);
		setInputText("");
	}

	const handleChange = (e) => {
		setInputText(e.target.value);
	}
	return (
		<div>
			<form onSubmit={handleSubmit}>
				<input 
					type="text" 
					placeholder="add a task" 
					className="taskAddInput"
					onChange={handleChange}
					value={inputText}
				/>
			</form>
		</div>
	)
}
