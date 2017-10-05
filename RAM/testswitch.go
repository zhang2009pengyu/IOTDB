package main

import (
    "fmt"
    "sort"
    "reflect"
    "encoding/json"
    "io/ioutil"
    "strconv"
    "path/filepath"
    "os"
)



type BinaryTree struct {
    Data   dataInfo
    Lchild *BinaryTree
    Rchild *BinaryTree
}

type dataInfo struct {
    time float64
    temperature float64
    ID int
}

//查询某点在某时刻的温度
func Find1(id int, time int){return}

    //查询某点在某时间段的温度
func Find2(id int, time []int){return    }

    //查询某区域在某时间点的温度
func find3(districtID []int, time int){return}

    //查询某区域在某时间内的温度
func find4(districtID []int, time []int){return}


//功能：判断结点是否为空结点
//输入：empty(dataInfo)
//输出：false/true
func empty(params interface{}) bool {  
    //初始化变量  
    var (  
        flag          bool = true  
        default_value reflect.Value  
    )  
  
    r := reflect.ValueOf(params)  
  
    //获取对应类型默认值  
    default_value = reflect.Zero(r.Type())  
    //由于params 接口类型 所以default_value也要获取对应接口类型的值 如果获取不为接口类型 一直为返回false  
    if !reflect.DeepEqual(r.Interface(), default_value.Interface()) {  
        flag = false  
    }  
    return flag  
}  

//功能：查询中位结点（用于kd树切分）
//输入：数组
//输出：数值
func FindMid(b []float64)float64{
    sort. Float64s(b)
    var pos int
    pos=len(b)/2
    var num float64
    num=b[pos]
    return num
}

func (tree *BinaryTree) PreOrder(tr *BinaryTree){
    if(tr==nil){ return }
    fmt.Println(tr)
    tr.PreOrder(tr.Lchild)
    tr.PreOrder(tr.Rchild)
}

//全局变量
var rootlist = make(map[int]dataInfo)
var datalist map[string]string

//功能：递归创建kd树
//输入：kd树根结点
//输出：kd树型结构
func (tree *BinaryTree) creat(data []dataInfo,slit int){
    var smallerarray []dataInfo
    var biggerarray []dataInfo
    var splitAttributeValues []float64

    //终止条件
    if(len(data)==0){ return }

    //选择切分属性
    if(slit==0){
        for i:=0;i<len(data);i++{
            splitAttributeValues= append(splitAttributeValues,data[i].time)
        }
    }else{
        for i:=0;i<len(data);i++{
            splitAttributeValues= append(splitAttributeValues,data[i].temperature)
        } 
    }

    //添加节点数据
    for i:=0;i<len(data);i++{
        if(slit==0 && data[i].time==FindMid(splitAttributeValues) && empty(tree.Data)==true){
            tree.Data=data[i]
            rootlist[data[i].ID]=data[i]
        }else if(slit==1 && data[i].temperature==FindMid(splitAttributeValues) && empty(tree.Data)==true){
        tree.Data=data[i]
        }else{ 
            if(slit==0){              
                if(data[i].time<FindMid(splitAttributeValues)){
                    smallerarray=append(smallerarray,data[i])
                }else{
                    biggerarray=append(biggerarray,data[i])
                }
            }else{
                if(data[i].temperature<FindMid(splitAttributeValues)){
                    smallerarray=append(smallerarray,data[i])
                }else{
                    biggerarray=append(biggerarray,data[i])
                }
            }
        }
    }

    //递归创建左右孩子结点
    var leftnew dataInfo
    tree.Lchild = NewBinTreeNode(leftnew)
    tree.Lchild.creat(smallerarray,(slit+1)%2)

    var rightnew dataInfo
    tree.Rchild = NewBinTreeNode(rightnew)
    tree.Rchild.creat(biggerarray,(slit+1)%2)
    }

func NewdataInfo(time float64, temperature float64, id int) dataInfo {
    return dataInfo{time:time,temperature:temperature,ID:id}
}

func NewBinTreeNode(e dataInfo) *BinaryTree {
    return &BinaryTree{Data: e, Rchild:nil,Lchild:nil}
}

//功能：读取json数据为map结构
//输入：json文件名称
//输出：map
func readFile(filename string) (map[string]string, error) {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("ReadFile: ", err.Error())
        return nil, err
    }

    if err := json.Unmarshal(bytes, &datalist); err != nil {
        fmt.Println("Unmarshal: ", err.Error())
        return nil, err
    }

    return datalist, nil
}
  
//功能：将string转化为int型
//输入：string类型
//输出：int类型
func changeInt(y string) int {
    //strconv.Atoi 就是将 string 类型 转成 int
    i, err := strconv.Atoi(y)
    if err != nil {
        panic(err)
    }
    return i
}

//功能：将string转化为float64型
//输入：string类型
//输出：float64类型
func change64(y string) float64 {
    //strconv.Atoi 就是将 string 类型 转成 int
    i, err := strconv.ParseFloat(y, 64)
        if err != nil {
        panic(err)
    }
    return i
}

type DataList []dataInfo

func main() {
filepath.Walk("/Users/pengpeng/Documents/go/testswitch/data/", //改为自己的data地址
func(path string, f os.FileInfo, err error) error {
    if f == nil {
        return err
    }
    if f.IsDir() {
        fmt.Println("dir:", path)
        return nil
    }
    fmt.Println("file:", path)

    var a *BinaryTree
    var root dataInfo
    var id int
    var pList []dataInfo

    //读入json数据
    datalist, err := readFile(path)
    if err != nil {
        fmt.Println("readFile: ", err.Error())
    }

    //查找id
    if v, ok := datalist["id"]; ok {
    id=changeInt(v)
    delete(datalist,"id")
    } else {
    fmt.Println("Key Not Found")
    }
    
    //将json原数据转化为结点数据结构
    for k, v := range datalist {
    fmt.Println(k, v)
    c:=NewdataInfo(change64(k),change64(v),id)
    pList=append(pList,c)
    }

    fmt.Println(datalist)  
    a = NewBinTreeNode(root)
    a.creat(pList,0)
    //fmt.Println(a) 

 
    a.PreOrder(a)

    return nil
})

fmt.Println(rootlist)
}
