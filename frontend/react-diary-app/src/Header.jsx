import { Box, Flex, Text } from "@chakra-ui/react";

export default function Header() {
	return (
		<Box>
			<Flex minH={'60px'}
				py={{ base: 2 }}
				px={{ base: 4 }}
				bgGradient='linear(to-l, #FF0080, #7928CA)'
				borderStyle={'solid'}
				align={'center'}>
				<Flex flex={{ base: 1 }} justify={{ base: 'center', md: 'start' }}>
		        	<Text as='b' fontSize='xl' color='white'>React Diary</Text>
		        </Flex>
			</Flex>
		</Box>
	);
}