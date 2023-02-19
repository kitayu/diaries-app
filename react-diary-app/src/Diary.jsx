import { Td, Tr } from '@chakra-ui/react'
import React from 'react'
import { Link } from 'react-router-dom'

export const Diary = ({no, title, id}) => {
  return (
    <Tr><Td>{no}</Td><Td><Link to={"/edit/" + id}>{title}</Link></Td></Tr>
  )
}
