// Generated by gen-tool
// DO NOT EDIT!
using System.Collections.Generic;
using System;
using ConfigProto.core;

namespace config
{

    public enum BasicItem_CommonQuality
    {

	yelow = 1,

	blue = 2,

	red = 3,

    }


    public class BasicItem_Common
    {
        public Dictionary<int, BasicItem_CommonInfo> Content;
    }
    public class BasicItem_CommonInfo
    {
	 
        public int Id;
	 
        public string Icon;
	 
        public int Quality;
	 
        public int Price;
	 
        public List<int> LimitNum;
	 
        public int Acauire;
	 
        public string ConsumeItem;
	 
        public int ConsumeCoin;
	 
        public int FormatIndex;
	 
        public BasicItem_CommonQuality Quality1;
	 
        public DateTime ShowTime;
	 
        public DateTime EndTime;
	 
        public Color TextColor;

    }
}

