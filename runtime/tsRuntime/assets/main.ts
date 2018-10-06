import {ConfigSerializer, ConfigTable} from "./ConfigTool/ConfigSerializer";

const {ccclass, property} = cc._decorator;

@ccclass
export default class NewClass extends cc.Component {

    @property(cc.Label)
    label: cc.Label = null;

    @property
    text: string = 'hello';

    // LIFE-CYCLE CALLBACKS:

    // onLoad () {}

    start () {
        this.LoadConfig("BasicItem_Common.bytes");
    }
    private LoadConfig(name){
        // load gun config
        var path = cc.url.raw('resources/config/' + name);
        cc.loader.load(path,function (err,res) {
            if(!err){
                let table : ConfigTable = ConfigSerializer.GetInstance().DeSerialize(res);
                if(table){
                    console.warn(table);
                    for(var i=0;i<table.FieldInfoList.length;i++){
                        console.warn(table.FieldInfoList[i].Type);
                        break;
                    }
                    console.warn(table.Type);
                }
            }else{
                console.error("error on load config by name " + path + " error : " + err);
            }
        }.bind(this));
    }
    // update (dt) {}
}
