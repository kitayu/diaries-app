import React, { useEffect, useState } from 'react'
import { Diary } from './Diary'

export const DiaryList = () => {
	const [diaries, setDiaries] = useState([]);

	useEffect(() => {
	  fetch("http://localhost:8080/api/diaries", {method: "GET"})
	  .then(res => res.json())
	  .then(data => {
		  setDiaries(data.diaries);
	  })
	}, [])
	return (
		diaries.map(diary => <Diary title={diary.title} key={diary.id}/>)
	);
}
