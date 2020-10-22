import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import DeleteIcon from '@material-ui/icons/Delete';
import {
  Content,
} from '@backstage/core';
import { EntPatient } from '../../api';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
  table: {
    minWidth: 650,
  },
  buttonRight: {
    marginLeft: theme.spacing(150),
    marginBottom: theme.spacing(2),
  },
  }),
);
 
export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [patients, setPatients] = useState<EntPatient[]>([]);
  const [loading, setLoading] = useState(true);
  

  // get patients
  useEffect(() => {
    const getPatients = async () => {
      const res = await api.listPatient({ limit: 10, offset: 0 });
      setLoading(true);
      setPatients(res);
      console.log(res);
    };
    getPatients();
  }, [loading]);
  
  // delete button
  const deletePatients = async (id: number) => {
    const res = await api.deletePatient({ id: id });
    setLoading(true);
  };

    // clear input form
    function clear() {
      setPatients([]);
    }
  
 
  // ui 
 return (
    <Content>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No</TableCell>
           <TableCell align="center">HN</TableCell>
           <TableCell align="center">ชื่อผู้ป่วย</TableCell>
           <TableCell align="center">เพศ</TableCell>
           <TableCell align="center">สิทธิการรักษา</TableCell>
           <TableCell align="center">เจ้าหน้าที่ที่ทำการบันทึก</TableCell>
         </TableRow>
       </TableHead>

       <TableBody>
         {patients.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.hn}</TableCell>
             <TableCell align="center">{item.patientName}</TableCell>
             <TableCell align="center">{item.edges?.gender?.genderName}</TableCell>
             <TableCell align="center">{item.edges?.rightoftreatment?.rightoftreatmentName}</TableCell>
             <TableCell align="center">{item.edges?.systemmember?.systemmemberName}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                   deletePatients(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
                 startIcon={<DeleteIcon/>}
                 href=""
               >
                 Delete
               </Button>
 
             </TableCell>

           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
);
}
