// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_DeschedulerPolicy, InType: reflect.TypeOf(&DeschedulerPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_DeschedulerStrategy, InType: reflect.TypeOf(&DeschedulerStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_NodeResourceUtilizationThresholds, InType: reflect.TypeOf(&NodeResourceUtilizationThresholds{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_StrategyParameters, InType: reflect.TypeOf(&StrategyParameters{})},
	)
}

// DeepCopy_v1alpha1_DeschedulerPolicy is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_DeschedulerPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeschedulerPolicy)
		out := out.(*DeschedulerPolicy)
		*out = *in
		if in.Strategies != nil {
			in, out := &in.Strategies, &out.Strategies
			*out = make(StrategyList)
			for key, val := range *in {
				newVal := new(DeschedulerStrategy)
				if err := DeepCopy_v1alpha1_DeschedulerStrategy(&val, newVal, c); err != nil {
					return err
				}
				(*out)[key] = *newVal
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_DeschedulerStrategy is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_DeschedulerStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeschedulerStrategy)
		out := out.(*DeschedulerStrategy)
		*out = *in
		if err := DeepCopy_v1alpha1_StrategyParameters(&in.Params, &out.Params, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_NodeResourceUtilizationThresholds is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_NodeResourceUtilizationThresholds(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NodeResourceUtilizationThresholds)
		out := out.(*NodeResourceUtilizationThresholds)
		*out = *in
		if in.Thresholds != nil {
			in, out := &in.Thresholds, &out.Thresholds
			*out = make(ResourceThresholds)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.TargetThresholds != nil {
			in, out := &in.TargetThresholds, &out.TargetThresholds
			*out = make(ResourceThresholds)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_StrategyParameters is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_StrategyParameters(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StrategyParameters)
		out := out.(*StrategyParameters)
		*out = *in
		if err := DeepCopy_v1alpha1_NodeResourceUtilizationThresholds(&in.NodeResourceUtilizationThresholds, &out.NodeResourceUtilizationThresholds, c); err != nil {
			return err
		}
		return nil
	}
}
