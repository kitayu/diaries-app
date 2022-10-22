import { Td, Tr } from '@chakra-ui/react'
import React from 'react'

export const Diary = ({no, title}) => {
  return (
    <Tr><Td>{no}</Td><Td>{title}</Td></Tr>
  )
}
