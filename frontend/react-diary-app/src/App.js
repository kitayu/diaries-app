import { ChakraProvider } from '@chakra-ui/react';
import './App.css';
import { DiaryList } from './DiaryList';

function App() {
  return (
    <ChakraProvider>
      <DiaryList />
    </ChakraProvider>
  );
}

export default App;
