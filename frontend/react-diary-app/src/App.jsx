import { ChakraProvider } from '@chakra-ui/react';
import './App.css';
import { DiaryList } from './DiaryList';
import Header from './Header';

function App() {
  return (
    <ChakraProvider>
      <Header />
      <DiaryList />
    </ChakraProvider>
  );
}

export default App;
