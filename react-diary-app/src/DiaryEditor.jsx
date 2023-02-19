import { Button, Container, Input, Textarea, VStack } from "@chakra-ui/react"
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

export const DiaryEditor = () => {
	const params = useParams();
	const navigate = useNavigate();
	const [diary, setDiary] = useState({});
	useEffect(() => {
		if (params.id) {
			fetch(`http://localhost:8080/api/diary/${params.id}`, {method: "GET"})
			.then(res => res.json())
			.then(data => {
				setDiary(data.diary);
			});
		}
	}, [params.id]);
	return (
		<VStack spacing={4}>
			<Container maxW='container.xl' mb={4}></Container>
				<Input value={diary?.title} onChange={e => {
					diary.title = e.target.value;
					setDiary({
						id: diary.id,
						title: e.target.value,
						description: diary.description,
					});
				}} placeholder='日記のタイトル' />
				<Textarea value={diary?.description}
					resize={"none"}
					placeholder='日記の内容を入力してください。'
					onChange={e => {
						diary.description = e.target.value;
						setDiary({
							id: diary.id,
							title: diary.title,
							description: e.target.value,
						});
					}} />
			<Container maxW='container.xl'></Container>
			<Container>
				<Button mr={4} onClick={() => navigate('/')}>キャンセル</Button>
				<Button colorScheme='purple' onClick={() => {
					if (params.id) {
						fetch(`http://localhost:8080/api/diary/${diary.id}`, {
							method: "PUT",
							headers: {
								"Content-Type": "application/json",
							},
							body: JSON.stringify({
								id: diary.id,
								title: diary.title,
								description: diary.description,
							}),
						})
						.then(res => res.json())
						.then(data => {
							setDiary(data.diary)
						})
					} else {
						fetch("http://localhost:8080/api/diary", {
							method: "POST",
							headers: {
								"Content-Type": "application/json",
							},
							body: JSON.stringify({
								title: diary.title,
								description: diary.description,
							}),
						})
						.then(res => res.json())
						.then(data => navigate('/edit/' + data.diary.id))
					}
				}}>保存</Button>
			</Container>
		</VStack>
	);
}

export default DiaryEditor;