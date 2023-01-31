// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package types

import (
	"fmt"
	"time"
)

type CloneRepoOptions struct {
	Timeout       time.Duration
	Mirror        bool
	Bare          bool
	Quiet         bool
	Branch        string
	Shared        bool
	NoCheckout    bool
	Depth         int
	Filter        string
	SkipTLSVerify bool
}

type SortOrder int

const (
	SortOrderDefault SortOrder = iota
	SortOrderAsc
	SortOrderDesc
)

type GitObjectType string

const (
	GitObjectTypeCommit GitObjectType = "commit"
	gitObjectTypeTree   GitObjectType = "tree"
	gitObjectTypeBlob   GitObjectType = "Blob"
	GitObjectTypeTag    GitObjectType = "tag"
)

func ParseGitObjectType(t string) (GitObjectType, error) {
	switch t {
	case string(GitObjectTypeCommit):
		return GitObjectTypeCommit, nil
	case string(gitObjectTypeBlob):
		return gitObjectTypeBlob, nil
	case string(gitObjectTypeTree):
		return gitObjectTypeTree, nil
	case string(GitObjectTypeTag):
		return GitObjectTypeTag, nil
	default:
		return gitObjectTypeBlob, fmt.Errorf("unknown git object type '%s'", t)
	}
}

// GitReferenceField represents the different fields available When listing references.
// For the full list, see https://git-scm.com/docs/git-for-each-ref#_field_names
type GitReferenceField string

const (
	GitReferenceFieldRefName     GitReferenceField = "refname"
	GitReferenceFieldObjectType  GitReferenceField = "objecttype"
	GitReferenceFieldObjectName  GitReferenceField = "objectname"
	GitReferenceFieldCreatorDate GitReferenceField = "creatordate"
)

func ParseGitReferenceField(f string) (GitReferenceField, error) {
	switch f {
	case string(GitReferenceFieldCreatorDate):
		return GitReferenceFieldCreatorDate, nil
	case string(GitReferenceFieldRefName):
		return GitReferenceFieldRefName, nil
	case string(GitReferenceFieldObjectName):
		return GitReferenceFieldObjectName, nil
	case string(GitReferenceFieldObjectType):
		return GitReferenceFieldObjectType, nil
	default:
		return GitReferenceFieldRefName, fmt.Errorf("unknown git reference field '%s'", f)
	}
}

type WalkInstruction int

const (
	WalkInstructionStop WalkInstruction = iota
	WalkInstructionHandle
	WalkInstructionSkip
)

type WalkReferencesEntry map[GitReferenceField]string

// TODO: can be generic (so other walk methods can use the same)
type WalkReferencesInstructor func(WalkReferencesEntry) (WalkInstruction, error)

// TODO: can be generic (so other walk methods can use the same)
type WalkReferencesHandler func(WalkReferencesEntry) error

type WalkReferencesOptions struct {
	// Patterns are the patterns used to pre-filter the references of the repo.
	// OPTIONAL. By default all references are walked.
	Patterns []string

	// Fields indicates the fields that are passed to the instructor & handler
	// OPTIONAL. Default fields are:
	// - GitReferenceFieldRefName
	// - GitReferenceFieldObjectName
	Fields []GitReferenceField

	// Instructor indicates on how to handle the reference.
	// OPTIONAL. By default all references are handled.
	// NOTE: once walkInstructionStop is returned, the walking stops.
	Instructor WalkReferencesInstructor

	// Sort indicates the field by which the references should be sorted.
	// OPTIONAL. By default GitReferenceFieldRefName is used.
	Sort GitReferenceField

	// Order indicates the Order (asc or desc) of the sorted output
	Order SortOrder

	// MaxWalkDistance is the maximum number of nodes that are iterated over before the walking stops.
	// OPTIONAL. A value of <= 0 will walk all references.
	// WARNING: Skipped elements count towards the walking distance
	MaxWalkDistance int32
}

type Commit struct {
	SHA       string
	Title     string
	Message   string
	Author    Signature
	Committer Signature
}

type Branch struct {
	Name   string
	SHA    string
	Commit *Commit
}

type Tag struct {
	Sha        string
	Name       string
	TargetSha  string
	TargetType GitObjectType
	Title      string
	Message    string
	Tagger     Signature
}

// Signature represents the Author or Committer information.
type Signature struct {
	Identity Identity
	// When is the timestamp of the Signature.
	When time.Time
}

type Identity struct {
	Name  string
	Email string
}

type CommitChangesOptions struct {
	Committer Signature
	Author    Signature
	Message   string
}

type PushOptions struct {
	Remote  string
	Branch  string
	Force   bool
	Mirror  bool
	Env     []string
	Timeout time.Duration
}

type TreeNodeWithCommit struct {
	TreeNode
	Commit *Commit
}

type TreeNode struct {
	NodeType TreeNodeType
	Mode     TreeNodeMode
	Sha      string
	Name     string
	Path     string
}

// TreeNodeType specifies the different types of nodes in a git tree.
// IMPORTANT: has to be consistent with rpc.TreeNodeType (proto).
type TreeNodeType int

const (
	TreeNodeTypeTree TreeNodeType = iota
	TreeNodeTypeBlob
	TreeNodeTypeCommit
)

// TreeNodeMode specifies the different modes of a node in a git tree.
// IMPORTANT: has to be consistent with rpc.TreeNodeMode (proto).
type TreeNodeMode int

const (
	TreeNodeModeFile TreeNodeMode = iota
	TreeNodeModeSymlink
	TreeNodeModeExec
	TreeNodeModeTree
	TreeNodeModeCommit
)

type Submodule struct {
	Name string
	URL  string
}

type Blob struct {
	Size int64
	// Content contains the content of the blob
	// NOTE: can be only partial Content - compare len(.Content) with .Size
	Content []byte
}

// CommitDivergenceRequest contains the refs for which the converging commits should be counted.
type CommitDivergenceRequest struct {
	// From is the ref from which the counting of the diverging commits starts.
	From string
	// To is the ref at which the counting of the diverging commits ends.
	To string
}

// CommitDivergence contains the information of the count of converging commits between two refs.
type CommitDivergence struct {
	// Ahead is the count of commits the 'From' ref is ahead of the 'To' ref.
	Ahead int32
	// Behind is the count of commits the 'From' ref is behind the 'To' ref.
	Behind int32
}

type PullRequest struct {
	BaseRepoPath string
	HeadRepoPath string

	BaseBranch string
	HeadBranch string
}

type DiffShortStat struct {
	Files     int
	Additions int
	Deletions int
}
