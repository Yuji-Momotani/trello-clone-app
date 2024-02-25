import React, { useContext } from 'react'
import { FaRegTimesCircle } from "react-icons/fa";
import { CardAndTaksContext } from '../TaskCards';
import taskCardUtil from "../../../util/taskCard";

export const TaskCardDeleteButton = ({taskCard}) => {
	const [, setTaskCardsList] = useContext(CardAndTaksContext);
	const {deleteCard} = taskCardUtil();
	const taskCardDelete = async(id) => {
		try {
			await deleteCard(taskCard.id);
			setTaskCardsList(prev => prev.filter(taskCard => taskCard.id !== id));
		} catch(e) {
			console.log(e);
			// 一旦無視
		}
	}
	return (
		<div>
			<button className="taskCardDeleteButton" onClick={() => taskCardDelete(taskCard.id)}>
				<FaRegTimesCircle />
			</button>
		</div>
	)
}
