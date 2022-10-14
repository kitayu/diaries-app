import { Container, Heading, Table, TableContainer, Tbody, Th, Thead, Tr, VStack } from '@chakra-ui/react';
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
		<VStack mt={4} mb={4}>
			<Container maxW='container.xl'>
				<Heading>日記一覧</Heading>
			</Container>
			<Container maxW='container.xl'>
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
			</Container>
		</VStack>
	);
}
