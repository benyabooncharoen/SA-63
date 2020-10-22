import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  Link,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import InputAdornment from '@material-ui/core/InputAdornment';
import PersonIcon from '@material-ui/icons/Person';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';

import { DefaultApi } from '../../api/apis';
import { EntGender } from '../../api/models/EntGender'; // import interface Gender
import { EntRightoftreatment } from '../../api/models/EntRightoftreatment'; // import interface RightOfTreatment
import { EntSystemmember } from '../../api/models/EntSystemmember'; // import interface Systemmember
import { EntPatient } from '../../api';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(1),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    select: {
      width: 500 ,
      marginLeft:7,
      marginRight:-7,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
  }),
);

/* interface CreatePatient {
  hn: string;
  patientName: string;
  gender: number;
  rightoftreatment: number;
  systemmember: number;
} */

export default function CreatePatientRecord() {
  const classes = useStyles();
  const api = new DefaultApi();
  
  const [patients, setPatient] = React.useState<EntPatient[]>([]);
 
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [rightoftreatments, setRightoftreatments] = React.useState<EntRightoftreatment[]>([]);
  const [systemmembers, setSystemmembers] = React.useState<EntSystemmember[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
  const [hn, SetHn] = useState(String);
  const [patientName, setPatientName] = useState(String);
  const [gender, setGender] = useState(Number);
  const [rightoftreatment, setRightoftreatment] = useState(Number);
  const [systemmember, setSystemmember] = useState(Number);

  useEffect(() => {
    const getGenders = async () => {
      const res = await api.listGender({ limit: 10, offset: 0 });
      setLoading(false);
      setGenders(res);
      console.log(res);
    };
    getGenders();
  
    const getRightoftreatments = async () => {
      const res = await api.listRightoftreatment({ limit: 10, offset: 0 });
      setLoading(false);
      setRightoftreatments(res);
      console.log(res);
    };
    getRightoftreatments();
  
    const getSystemmembers = async () => {
      const res = await api.listSystemmember({ limit: 10, offset: 0 });
      setLoading(false);
      setSystemmembers(res);
      console.log(res);
    };
    getSystemmembers();
  
  }, [loading]);
  
  
  const getPatient = async () => {
    const res = await api.listPatient({ limit: 10, offset: 0 });
    setPatient(res);
  };
  
  const handlehnChange = (event: any) => {
    SetHn(event.target.value as string);
  };
  
  const handlepatientNameChange = (event: any) => {
    setPatientName(event.target.value as string);
  };
  
  const GenderhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setGender(event.target.value as number);
  };
  
  const RightoftreatmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRightoftreatment(event.target.value as number);
  };
  
  const SystemmemberhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSystemmember(event.target.value as number);
  };
  
  // create patient
  const CreatePatient = async () => {
    const patient = {
      hn : hn,
      patientName : patientName,
      gender : gender,
      rightoftreatment : rightoftreatment,
      systemmember : systemmember,
    };
    console.log(patient);
    const res: any = await api.createPatient({ patient : patient });
    console.log("bruhhhhhhhhh");
    setStatus(true);
    
  };

  return (
    <Page theme={pageTheme.tool}>

      <Header
        title={`ทะเบียนผู้ป่วยใน`}> 
      </Header>
      <Content>
        <ContentHeader title="ลงทะเบียนผู้ป่วย">
        <div>
            <Link component={RouterLink} to="/">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              Exit
            </Button>
          </Link>
          </div>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              //fullWidth
              //className={classes.margin}
              variant="outlined"
            >
               <div className={classes.paper}><strong>หมายเลข HN</strong></div>
              <TextField className={classes.textField}
               style={{ width: 400 ,marginLeft:9,marginRight:-9}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <PersonIcon />
                  </InputAdornment>
                ),
              }}
                id="hn"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={hn}
                onChange={handlehnChange}
              />

            <div className={classes.paper}><strong>ชื่อผู้ป่วย</strong></div>
              <TextField className={classes.textField}
              style={{ width: 400 ,marginLeft:7,marginRight:-7}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <PersonIcon />
                  </InputAdornment>
                ),
              }}
                id="patientName"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={patientName}
                onChange={handlepatientNameChange}
              />
            
            <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <Typography gutterBottom  align="left">
                เพศ :   
              <Typography variant="body1" gutterBottom> 
              <Select
                labelId="gender"
                id="gender"
                value={gender}
                onChange={GenderhandleChange}
                style={{ width: 400 }}
              >
                {genders.map((item: EntGender) => (
                  <MenuItem value={item.id}>{item.genderName}</MenuItem>
                ))}
              </Select>
              </Typography>
                </Typography>
            </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <Typography gutterBottom  align="left">
                  สิทธิการรักษา : 
                <Typography variant="body1" gutterBottom> 
                <Select
                   labelId="rightoftreatment"
                   id="rightoftreatment"
                   value={rightoftreatment}
                   onChange={RightoftreatmenthandleChange}
                   style={{ width: 400 }}
                >
                  {rightoftreatments.map((item: EntRightoftreatment) => (
                  <MenuItem value={item.id}>{item.rightoftreatmentName}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography>
              </FormControl>
              </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <Typography gutterBottom  align="left">
                  เจ้าหน้าที่ที่ทำการบันทึก : 
                <Typography variant="body1" gutterBottom> 
                <Select
                   labelId="systemmember"
                   id="systemmember"
                   value={systemmember}
                   onChange={SystemmemberhandleChange}
                   style={{ width: 400 }}
                >
                  {systemmembers.map((item: EntSystemmember) => (
                  <MenuItem value={item.id}>{item.systemmemberName}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography>
              </FormControl>
            </div>

            <div className={classes.margin}>
                <Typography variant="h6" gutterBottom  align="center">
                  <Button
                    onClick={() => {
                      CreatePatient();
                    }}
                      variant="contained"
                      color="primary"
                  >
                    Submit
                  </Button>
                </Typography>
              </div>
              </FormControl>

          </form>
        </div>
    </Content>
</Page>
);
} 