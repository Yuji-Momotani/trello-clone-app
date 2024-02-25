# Trelloクローンアプリ
## 概要
フロントエンド参考講座：[ReactでTrelloクローンアプリケーションを作ってReactをマスターしよう！](https://www.udemy.com/course/react-trello-development/)

## 学び
### react-beautiful-dndでドラッグ可能なUI

[公式](https://github.com/atlassian/react-beautiful-dnd#documentation-)

- 注意点
あまりメンテナンスが頻繁にされてなさそうで、React18系では不具合が起きてるっぽい。
一旦、Strictモードを解除して問題がないか確認中

- インストール
```sh
npm install react-beautiful-dnd
```

- react-beautiful-dndのお作法

```js
function App() {
	const onDragEnd = (result) => {
		console.log(result);
	}
  return (
    <div className="dragDropArea">
			<DragDropContext onDragEnd={onDragEnd}>
				<Droppable droppableId="droppable">
					{(provided) => (
						<div {...provided.droppableProps} ref={provided.innerRef}>
							<Draggable draggableId="item0" index={0}>
								{(provided) => (
									<div className="item"
										ref={provided.innerRef} 
										{...provided.draggableProps} 
										{...provided.dragHandleProps}
									>Items0</div>
								)}
							</Draggable>
							<Draggable draggableId="item1" index={1}>
								{(provided) => (
									<div className="item"
										ref={provided.innerRef} 
										{...provided.draggableProps} 
										{...provided.dragHandleProps}
									>Items1</div>
								)}
							</Draggable>
							<Draggable draggableId="item2" index={2}>
								{(provided) => (
									<div className="item"
										ref={provided.innerRef} 
										{...provided.draggableProps} 
										{...provided.dragHandleProps}
									>Items2</div>
								)}
							</Draggable>
							{provided.placeholder}
						</div>
					)}
				</Droppable>
			</DragDropContext>
    </div>
  );
}
```

- DragDropContext、Droppable、Draggableの関係性は以下

![Alt text](/doc/images/react-beautiful-dnd-image.png)

- onDragEnd
	- ドラッグの操作が終わった時に発火する関数をセットできる。
Droppable、Draggableの子要素は関数でないといけない。

この状態では、まだドラッグ後の処理が記述できていないので、要素の入れ替わりはできない。

- 要素の入れ替え処理を実装

```js
function App() {
	const [items, setItems] = useState([
		{id:0, text: "item0"},
		{id:1, text: "item1"},
		{id:2, text: "item2"},
	]);
	const onDragEnd = (result) => {
		setItems(item => {
			items.splice(result.destination.index, 0, item.splice(result.source.index, 1)[0]);
			return items;
		})
	}
  return (
    <div className="dragDropArea">
			<DragDropContext onDragEnd={onDragEnd}>
				<Droppable droppableId="droppable">
					{(provided) => (
						<div {...provided.droppableProps} ref={provided.innerRef}>
							{
								items.map((item, index) => (
									<Draggable draggableId={item.text} index={index} key={item.id}>
										{(provided) => (
											<div className="item"
												ref={provided.innerRef} 
												{...provided.draggableProps} 
												{...provided.dragHandleProps}
											>{item.text}</div>
										)}
									</Draggable>
								))
							}
							{provided.placeholder}
						</div>
					)}
				</Droppable>
			</DragDropContext>
    </div>
  );
}
```

配列の要素の入れ替えはArray.spliceで実行
例：
```js
const array = ["test0", "test1", "test2"];

// 要素の削除
const remove = array.splice(0, 1);
// array => ["test1", "test2"]
// remove => ["test0"]
// spliceの第一引数は配列のindex番号。第二引数は削除する要素数

//ここで要素の入れ替わりが完成
array.splice(1, 0, remove[0]);
// array => ["test1", "test0","test2"]
// spliceの第三引数は追加する要素
```
このようにArray.spliceを使用すれば配列に対して、要素の追加、削除が可能

### react-beautiful-dndの代替ツール

react-beautiful-dndは今後、あまり積極的にメンテナンスされないようだ。
代替ツールとしては以下があるっぽい。
- [react-dnd](https://github.com/react-dnd/react-dnd)
- [dnd-kit](https://dndkit.com/)

dnd-kitは最近伸びてきているっぽいし、公式ページも綺麗で今後使うには良さそう！


