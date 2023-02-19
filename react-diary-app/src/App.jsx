import { ChakraProvider } from '@chakra-ui/react';
import Header from './Header';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import React, { Suspense } from 'react';

const DiaryList = React.lazy(() => import('./DiaryList'));
const DiaryEditor = React.lazy(() => import('./DiaryEditor'));

function App() {
  return (
    <Router>
      <Suspense fallback={<div>Loading...</div>}>
        <ChakraProvider>
          <Header />
          <Routes>
            <Route path="/" element={<DiaryList />} />
            <Route path="/edit/:id" element={<DiaryEditor />} />
            <Route path="/newedit" element={<DiaryEditor />} />
          </Routes>
        </ChakraProvider>
      </Suspense>
      </Router>
  );
}

export default App;
