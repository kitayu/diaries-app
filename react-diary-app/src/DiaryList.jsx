import { Button, Container, Flex, Heading, Spacer, Table, TableContainer, Tbody, Th, Thead, Tr, VStack } from '@chakra-ui/react';
import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { Diary } from './Diary'

export const DiaryList = () => {
	const [diaries, setDiaries] = useState([]);
	const navigate = useNavigate();
	useEffect(() => {
	  fetch("http://localhost:8080/api/diaries", {method: "GET"})
	  .then(res => res.json())
	  .then(data => {
		  setDiaries(data.diaries);
	  })
	}, [])
	return (
		<>
		<VStack mt={4} mb={4}>
			<Container maxW='container.xl'>
				<Flex>
					<Heading>日記一覧</Heading>
					<Spacer />
					<Button
						colorScheme='purple'
						onClick={() => navigate('/newedit')}
					>新しく日記を書く</Button>
				</Flex>
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
							{diaries.map((diary, index) => <Diary no={index+1} title={diary.title} id={diary.id} key={diary.id}/>)}
						</Tbody>
					</Table>
				</TableContainer>
			</Container>
		</VStack>
		</>
	);
}

export default DiaryList;