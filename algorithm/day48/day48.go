package main

// 二叉树的下一个结点

/*
给定一个二叉树其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。
注意，树中的结点不仅包含左右子结点，同时包含指向父结点的next指针。下图为一棵有9个节点的二叉树。
树中从父节点指向子节点的指针用实线表示，从子节点指向父节点的用虚线表示
*/

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

// 像刚才那种情况怎么处理呢
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {

	if pNode.Right != nil {
		tem := pNode.Right
		for tem.Left != nil {
			tem = tem.Left
		}
		return tem
	}
	if pNode.Next != nil && pNode == pNode.Next.Left {
		return pNode.Next
	} else if pNode.Next != nil && pNode == pNode.Next.Right {
		tem := pNode
		for tem.Next != nil && tem == tem.Next.Right {
			tem = tem.Next
		}
		return tem.Next
	}

	return nil
}

func main() {

}

/*
偷偷写封给过去的信
信中是一眨眼的光阴
早上总是是是睡不醒
窗外云高风清
午后的蝉鸣清脆的很好听
孩子们在黄昏的落日里做游戏
小小时光里充满了欢声与笑语
现在的眼前陌生而熟悉

偷偷写份给过去的信
信中写满沉甸甸的歉意
允许对过去的稚气说句对不起
无意伤害过的你
还有深夜的啜泣
现在都藏在了涓涓心底


偷偷写份给过去的信
记录星光里的脚步不曾停
长野与星空 或倾盆与雷鸣都是一路最好的风景
来过去过都是最好的天意
希望在某个日子漫天繁星
慢慢说给你听

*/
