package main

type Block struct {
  Index string `json:"index"`
  Timestamp string `json:"timestamp"`
  Proof long `json:"proof"`
  PreviousBlockHash  `json:" `json:"previousBlockHash"`
  ....
}
