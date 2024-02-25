import React, { createContext, useEffect, useState } from 'react'
import { TaskCard } from "./TaskCard"
import { AddTaskCardButton } from "./button/AddTaskCardButton"
import { DragDropContext, Droppable } from "react-beautiful-dnd";
import taskUtil from "../../util/task";
import taskCardUtil from "../../util/taskCard";

const {updateCard} = taskCardUtil();

const reorder = (taskCardsList, fromIndex, toIndex) => {
	const oldTaskCardsList = taskCardsList.map(task => task);
	const removeFromTask = taskCardsList.splice(fromIndex, 1)[0];

	taskCardsList.splice(toIndex, 0, removeFromTask);

	const newTaskCardsList = taskCardsList.map((taskCard, index) => {
		taskCard.sort_no = index;
		return taskCard;
	});

	try {
		const updateTaskCardList = newTaskCardsList.filter((taskCard,index) => {
			return oldTaskCardsList[index].id !== taskCard.id; // sort_noの変更があるもののみに絞り込み
		});
		updateTaskCardList.forEach(async(taskCard) => {
			await updateCard(taskCard.id, taskCard.title, taskCard.sort_no);
		});
	} catch (e) {
		console.error(e);
		//　一旦何もしない
	}
	return newTaskCardsList;
}

export const CardAndTaksContext = createContext();

export const TaskCards = () => {
	const {getTaskCardAndTask} = taskUtil();
	const [taskCardsList, setTaskCardsList] = useState([]);

	useEffect(() => {
		const getInitData = async() => {
			try {
				const { data } = await getTaskCardAndTask();
				console.log(data);
				setTaskCardsList(() => {
					const initData = data.map(card => {
						return {
							id: card.id,
							sort_no: card.sort_no,
							title: card.title,
							draggableId: `item-${card.id}`,
							tasks: card.tasks ? card.tasks : [],
						}
					});
					return initData;
				});
			} catch(e) {
				// 一旦何もしない
			}
		}
		getInitData();
	},[]);
	const handleDragEnd = async(result) => {
		const data = await reorder(taskCardsList, result.source.index, result.destination.index);
		setTaskCardsList(data);
	}
	return (
		<DragDropContext onDragEnd={handleDragEnd}>
			<Droppable droppableId="droppable" direction="horizontal">
				{(provided) => (
					<CardAndTaksContext.Provider value={[taskCardsList, setTaskCardsList]}>
						<div className="taskCardArea" {...provided.droppableProps} ref={provided.innerRef}>
							{
								taskCardsList.map((taskCard, index) => (
									<TaskCard key={taskCard.id} taskCard={taskCard} draggableIndex={index} />
								))
							}
							{provided.placeholder}
							<AddTaskCardButton 
								// setTaskCardsList={setTaskCardsList}
							/>
						</div>
					</CardAndTaksContext.Provider>
				)}
			</Droppable>
		</DragDropContext>
	)
}
