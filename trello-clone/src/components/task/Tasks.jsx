import React, { useContext } from 'react'
import { Task } from "./Task"
import { DragDropContext, Droppable } from "react-beautiful-dnd"
import { TaskListContext } from "./TaskCard";
import taskUtil from "../../util/task";

const {updateTask} = taskUtil();

const reorder = async(taskList, fromIndex, toIndex) => {
	const oldTaskList = taskList.map(task => task);
	const removeFromTask = taskList.splice(fromIndex, 1)[0];

	taskList.splice(toIndex, 0, removeFromTask);

	const newTaskList = taskList.map((task, index) => {
		task.sort_no = index;
		return task;
	});

	try {
		const updateTaskList = newTaskList.filter((task,index) => {
			return oldTaskList[index].id !== task.id; // sort_noの変更があるもののみに絞り込み
		});
		updateTaskList.forEach(async(task) => {
			await updateTask(task.id, task.contetn, task.sort_no);
		});
	} catch (e) {
		console.error(e);
		//　一旦何もしない
	}
	return newTaskList;
}

export const Tasks = () => {

	const [taskList, setTaskList] = useContext(TaskListContext);

	const handleDragEnd = async(result) => {
		// タスクを並び変える。
		const data = await reorder(taskList, result.source.index, result.destination.index);
		setTaskList(data);
	};

	return (
		<div>
			<DragDropContext onDragEnd={handleDragEnd}>
				<Droppable droppableId="doroppable">
					{(provided) => (
						<div {...provided.droppableProps} ref={provided.innerRef}>
							{taskList.map((task,index) => (
								<div key={task.id}>
									<Task task={task} draggableIndex={index} />
								</div>
							))}
							{provided.placeholder}
						</div>
					)}
				</Droppable>
			</DragDropContext>
		</div>
	)
}
