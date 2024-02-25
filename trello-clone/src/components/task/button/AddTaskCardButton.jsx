import React, { useContext } from 'react'
import { CardAndTaksContext } from '../TaskCards';
import taskCardUtil from "../../../util/taskCard";

const DEFAULT_REGIST_CARD_TITLE = "Today";

export const AddTaskCardButton = () => {
	const [taskCardList, setTaskCardsList] = useContext(CardAndTaksContext);
	const {registCard} = taskCardUtil();
	const addTaskCard = async () => {
		try {
			let maxSortNo = -1;
			if (taskCardList.length === 1) {
				maxSortNo = taskCardList[0].sort_no;
			} else if (taskCardList.length > 1) {
				maxSortNo = taskCardList.reduce((x, y) => x.sort_no > y.sort_no ? x.sort_no: y.sort_no, 0);
			}
			const {data} = await registCard(DEFAULT_REGIST_CARD_TITLE, maxSortNo+1);
			setTaskCardsList(prev => {
				return [...prev, {
					id: data.id,
					sort_no: data.sort_no,
					title: data.title,
					draggableId: `item-${data.id}`,
					tasks: [],
				}]
			});
		} catch(e) {
			// 一旦何もしない
		}
	}
	return (
		<div className="addTaskCardButtonArea">
			<button className="addTaskCardButton" onClick={addTaskCard}>+</button>
		</div>
	)
}
