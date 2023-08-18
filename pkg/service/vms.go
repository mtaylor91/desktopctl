package service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sTypes "k8s.io/apimachinery/pkg/types"
)

var patchVMStart = []byte(`[
	{
		"op": "replace",
		"path": "/spec/running",
		"value": true
	}
]`)

var patchVMStop = []byte(`[
	{
		"op": "replace",
		"path": "/spec/running",
		"value": false
	}
]`)

// getVM is a handler for GET /vms/{name}
func (s *Service) getVM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vm, err := s.client.KubevirtV1().
		VirtualMachines(s.namespace).
		Get(r.Context(), vars["name"], metav1.GetOptions{})
	if err != nil && errors.IsNotFound(err) {
		http.NotFound(w, r)
		return
	} else if err != nil {
		panic(err)
	}

	// Return the VM as JSON
	err = json.NewEncoder(w).Encode(vm)
	if err != nil {
		panic(err)
	}
}

// listVMs is a handler for GET /vms
func (s *Service) listVMs(w http.ResponseWriter, r *http.Request) {
	vms, err := s.client.KubevirtV1().
		VirtualMachines(s.namespace).
		List(r.Context(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// Return the list of VMs as JSON
	err = json.NewEncoder(w).Encode(vms)
	if err != nil {
		panic(err)
	}
}

// startVM is a handler for POST /vms/{name}/start
func (s *Service) startVM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vm, err := s.client.KubevirtV1().VirtualMachines(s.namespace).
		Patch(r.Context(), vars["name"], k8sTypes.JSONPatchType,
			patchVMStart, metav1.PatchOptions{})
	if err != nil && errors.IsNotFound(err) {
		http.NotFound(w, r)
		return
	} else if err != nil {
		panic(err)
	}

	// Return the VM as JSON
	err = json.NewEncoder(w).Encode(vm)
	if err != nil {
		panic(err)
	}
}

// stopVM is a handler for POST /vms/{name}/stop
func (s *Service) stopVM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vm, err := s.client.KubevirtV1().VirtualMachines(s.namespace).
		Patch(r.Context(), vars["name"], k8sTypes.JSONPatchType,
			patchVMStop, metav1.PatchOptions{})
	if err != nil && errors.IsNotFound(err) {
		http.NotFound(w, r)
		return
	} else if err != nil {
		panic(err)
	}

	// Return the VM as JSON
	err = json.NewEncoder(w).Encode(vm)
	if err != nil {
		panic(err)
	}
}
