import { Table, TableContainer, Tbody, Td, Th, Thead, Tr } from '@chakra-ui/react';
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
		<TableContainer>
			<Table variant='simple'>
				<Thead>
					<Tr>
						<Th>No.</Th>
						<Th>Title</Th>
					</Tr>
				</Thead>
				<Tbody>
					{diaries.map((diary, index) => <Diary no={index+1} title={diary.title} key={diary.id}/>)}
				</Tbody>
			</Table>
		</TableContainer>
	);
}
