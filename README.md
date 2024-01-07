# Go Gale-Shapley

An implementation of the Gale-Shapley algorithm in Golang, with minimal depdendencies. 

In this implementation two groups (generally people) are matched together in a stable way, assuming the group sizes are equal or near equal. A list of enumerated preferences for each member of the group must be provided.

## Usage

Follow these steps from `main.go`, which calls the `stableMatching()` function defined in `matching.go`.

First, define preferences: before running the program, define the preference lists for both groups (defined as Group A and Group B). These preferences are arrays of integers, where each array represents the preference list of a member of a group. What these enumerations actually mean are up to you. For example:

```go
groupAPref := [][]int{{3, 1, 2, 0}, {1, 0, 2, 3}, {0, 1, 2, 3}, {0, 1, 2, 3}}
groupBPref := [][]int{{0, 1, 2, 3}, {0, 1, 2, 3}, {0, 1, 2, 3}, {0, 1, 2, 3}}
```

In this example, each member of Group A and Group B has ranked the members of the opposite group in order of preference, with 0 being the most preferred.

Next, run the program like

```bash
go run main.go
```

This will output the stable matches between members of Group A and Group B.

## Interpreting the Output

The output will be a series of lines indicating the matching pairs, with each line showing a member of Group B and their matched partner in Group A.

## Customization

You can modify `main.go` to change the preferences or the number of members in each group. The `stableMatching` function will handle the matching process based on the provided preferences.
