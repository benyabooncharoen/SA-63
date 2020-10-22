import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
import PersonAddRoundedIcon from '@material-ui/icons/PersonAddRounded';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ทะเบียนผู้ป่วยใน' }; 
 
 return (
   <Page theme={pageTheme.tool}>
     <Header
       title={`${profile.givenName || 'to Backstage'}`}
     ></Header>
     <Content>
       <ContentHeader title="ข้อมูลผู้ป่วย">
         <Link component={RouterLink} to="/user">
         <div>&nbsp;&nbsp;&nbsp;</div>
            <Button variant="contained" color="primary" href="/recorduser" startIcon={<PersonAddRoundedIcon />}> ADD PATIENT</Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
