import dateincrement from 'k6/x/atb/dateincrement';
import numincrement from 'k6/x/atb/numericincrement';

const dincrement =  dateincrement.New("2006-01-02",1)
const numincrement = numincrement.New(1,1)

export default function(){
    console.log(dincrement.Get())
    console.log(numincrement.Get())
}