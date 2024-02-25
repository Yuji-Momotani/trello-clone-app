import { FaRegTrashAlt } from "react-icons/fa";
import React, { useContext } from 'react'
import { Draggable } from "react-beautiful-dnd";
import { TaskCardContext, TaskListContext } from "./TaskCard";
import taskUtil from "../../util/task";

export const Task = ({task, draggableIndex}) => {

	const [, setTaskList] = useContext(TaskListContext);
	const taskCard = useContext(TaskCardContext);
	const {deleteTask} = taskUtil();

	const handleDelete = async(id) => {
		try {
			await deleteTask(id, taskCard.id, task.sort_no);
			setTaskList(prev => {
				const newTaskList = prev.filter(task => {
					return task.id !== id;
				});
				return newTaskList;
			});
		} catch(e) {
			console.log(e);
			// 一旦何もしない
		}
	}
	return (
		<Draggable index={draggableIndex} draggableId={`task-${task.id}`}>
			{(provided) => (
				<div className="taskBox" 
					key={task.id} 
					ref={provided.innerRef} 
					{...provided.draggableProps} 
					{...provided.dragHandleProps}
				>
					<p className="taskContent">{task.contetn}</p>
					<button className="taskTrashButton" onClick={() => handleDelete(task.id)}>
						<FaRegTrashAlt />
					</button>
				</div>
			)}
		</Draggable>
	)
}
